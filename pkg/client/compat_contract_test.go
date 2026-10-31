// Copyright © 2026 KubeCub. All rights reserved.
// Licensed under the MIT License (the "License").

package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	upstream "github.com/PullRequestInc/go-gpt3"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type capturedRequest struct {
	method, path, authorization, contentType string
	body                                     map[string]interface{}
	err                                      error
}

func firstCompletion(result *CompletionResponse, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if result == nil || len(result.Choices) != 1 {
		return "", errors.New("missing synthetic completion")
	}
	return result.Choices[0].Text, nil
}

func TestClientDelegatesAllLegacyMethods(t *testing.T) {
	ctx := context.Background()
	completion := CompletionRequest{Prompt: []string{"synthetic prompt"}, MaxTokens: IntPtr(10)}
	chat := ChatCompletionRequest{Model: "gpt-3.5-turbo", Messages: []ChatCompletionRequestMessage{{Role: "user", Content: "synthetic prompt"}}}
	search := SearchRequest{Documents: []string{"synthetic document"}, Query: "synthetic query"}
	completionJSON := `{"id":"synthetic-completion","choices":[{"index":0,"text":"synthetic reply"}]}`
	completionSSE := "data: " + completionJSON + "\n\ndata: [DONE]\n\n"
	searchJSON := `{"data":[{"document":0,"object":"search_result","score":0.5}]}`
	cases := []struct {
		name, method, path, response, want string
		body                               map[string]interface{}
		call                               func(Client) (string, error)
	}{
		{"engines", "GET", "/engines", `{"data":[{"id":"synthetic-engine"}]}`, "synthetic-engine", nil, func(c Client) (string, error) {
			result, err := c.Engines(ctx)
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Data) != 1 {
				return "", errors.New("missing synthetic engine list")
			}
			return result.Data[0].ID, nil
		}},
		{"engine", "GET", "/engines/synthetic-engine", `{"id":"synthetic-engine"}`, "synthetic-engine", nil, func(c Client) (string, error) {
			result, err := c.Engine(ctx, "synthetic-engine")
			if err != nil {
				return "", err
			}
			if result == nil {
				return "", errors.New("missing synthetic engine")
			}
			return result.ID, nil
		}},
		{"chat", "POST", "/chat/completions", `{"choices":[{"message":{"role":"assistant","content":"synthetic reply"}}]}`, "synthetic reply", map[string]interface{}{"model": "gpt-3.5-turbo", "stream": nil}, func(c Client) (string, error) {
			result, err := c.ChatCompletion(ctx, chat)
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Choices) != 1 {
				return "", errors.New("missing synthetic chat")
			}
			return result.Choices[0].Message.Content, nil
		}},
		{"chat stream", "POST", "/chat/completions", "data: {\"choices\":[{\"delta\":{\"content\":\"synthetic piece\"}}]}\n\ndata: [DONE]\n\n", "synthetic piece", map[string]interface{}{"model": "gpt-3.5-turbo", "stream": true}, func(c Client) (string, error) {
			var pieces []string
			err := c.ChatCompletionStream(ctx, chat, func(result *ChatCompletionStreamResponse) {
				for _, choice := range result.Choices {
					pieces = append(pieces, choice.Delta.Content)
				}
			})
			return strings.Join(pieces, ""), err
		}},
		{"completion", "POST", "/engines/davinci/completions", completionJSON, "synthetic reply", map[string]interface{}{"prompt": []interface{}{"synthetic prompt"}, "max_tokens": float64(10), "stream": nil}, func(c Client) (string, error) { return firstCompletion(c.Completion(ctx, completion)) }},
		{"completion with engine", "POST", "/engines/synthetic-engine/completions", completionJSON, "synthetic reply", map[string]interface{}{"max_tokens": float64(10), "stream": nil}, func(c Client) (string, error) {
			return firstCompletion(c.CompletionWithEngine(ctx, "synthetic-engine", completion))
		}},
		{"completion stream", "POST", "/engines/davinci/completions", completionSSE, "synthetic reply", map[string]interface{}{"stream": true}, func(c Client) (string, error) {
			var pieces []string
			err := c.CompletionStream(ctx, completion, func(result *CompletionResponse) {
				for _, choice := range result.Choices {
					pieces = append(pieces, choice.Text)
				}
			})
			return strings.Join(pieces, ""), err
		}},
		{"completion stream with engine", "POST", "/engines/synthetic-engine/completions", completionSSE, "synthetic reply", map[string]interface{}{"stream": true}, func(c Client) (string, error) {
			var pieces []string
			err := c.CompletionStreamWithEngine(ctx, "synthetic-engine", completion, func(result *CompletionResponse) {
				for _, choice := range result.Choices {
					pieces = append(pieces, choice.Text)
				}
			})
			return strings.Join(pieces, ""), err
		}},
		{"retired edits contract only", "POST", "/edits", `{"choices":[{"text":"synthetic edit","index":0}]}`, "synthetic edit", map[string]interface{}{"model": "text-davinci-edit-001", "input": "synthetic input", "instruction": "synthetic instruction"}, func(c Client) (string, error) {
			result, err := c.Edits(ctx, EditsRequest{Model: "text-davinci-edit-001", Input: "synthetic input", Instruction: "synthetic instruction"})
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Choices) != 1 {
				return "", errors.New("missing synthetic edit")
			}
			return result.Choices[0].Text, nil
		}},
		{"search", "POST", "/engines/davinci/search", searchJSON, "0.5", map[string]interface{}{"query": "synthetic query"}, func(c Client) (string, error) {
			result, err := c.Search(ctx, search)
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Data) != 1 {
				return "", errors.New("missing synthetic search")
			}
			return fmt.Sprint(result.Data[0].Score), nil
		}},
		{"search with engine", "POST", "/engines/synthetic-engine/search", searchJSON, "0.5", map[string]interface{}{"query": "synthetic query"}, func(c Client) (string, error) {
			result, err := c.SearchWithEngine(ctx, "synthetic-engine", search)
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Data) != 1 {
				return "", errors.New("missing synthetic search")
			}
			return fmt.Sprint(result.Data[0].Score), nil
		}},
		{"embeddings", "POST", "/embeddings", `{"data":[{"embedding":[0.5],"index":0}]}`, "0.5", map[string]interface{}{"model": "text-embedding-ada-002", "input": []interface{}{"synthetic input"}}, func(c Client) (string, error) {
			result, err := c.Embeddings(ctx, EmbeddingsRequest{Model: "text-embedding-ada-002", Input: []string{"synthetic input"}})
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Data) != 1 || len(result.Data[0].Embedding) != 1 {
				return "", errors.New("missing synthetic embedding")
			}
			return fmt.Sprint(result.Data[0].Embedding[0]), nil
		}},
		{"moderation", "POST", "/moderations", `{"results":[{"flagged":false}]}`, "false", map[string]interface{}{"input": "synthetic input", "model": "text-moderation-latest"}, func(c Client) (string, error) {
			result, err := c.Moderation(ctx, ModerationRequest{Input: "synthetic input", Model: "text-moderation-latest"})
			if err != nil {
				return "", err
			}
			if result == nil || len(result.Results) != 1 {
				return "", errors.New("missing synthetic moderation")
			}
			return fmt.Sprint(result.Results[0].Flagged), nil
		}},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			requests := make(chan capturedRequest, 2)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				raw, err := io.ReadAll(r.Body)
				var body map[string]interface{}
				if err == nil && len(raw) > 0 {
					err = json.Unmarshal(raw, &body)
				}
				requests <- capturedRequest{r.Method, r.URL.Path, r.Header.Get("Authorization"), r.Header.Get("Content-Type"), body, err}
				if strings.Contains(test.response, "data: ") {
					w.Header().Set("Content-Type", "text/event-stream")
				} else {
					w.Header().Set("Content-Type", "application/json")
				}
				_, _ = io.WriteString(w, test.response)
			}))
			t.Cleanup(server.Close)
			base := http.DefaultTransport.(*http.Transport).Clone()
			base.Proxy = nil
			t.Cleanup(base.CloseIdleConnections)
			localURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatal(err)
			}
			guarded := roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Host != localURL.Host {
					return nil, errors.New("non-fixture destination rejected")
				}
				return base.RoundTrip(r)
			})
			c := NewClient("synthetic-key", upstream.WithBaseURL(server.URL), upstream.WithHTTPClient(&http.Client{Transport: guarded, Timeout: 2 * time.Second}))
			got, err := test.call(c)
			if err != nil || got != test.want {
				t.Fatalf("actual client method returned %q, %v", got, err)
			}
			select {
			case request := <-requests:
				if request.err != nil || request.method != test.method || request.path != test.path {
					t.Fatalf("request contract changed: %s %s, %v", request.method, request.path, request.err)
				}
				if request.authorization != "Bearer synthetic-key" || request.contentType != "application/json" {
					t.Fatal("request authentication/content-type contract changed")
				}
				for key, want := range test.body {
					if want == nil {
						if _, present := request.body[key]; present {
							t.Fatalf("omitempty JSON field %s unexpectedly became present", key)
						}
						continue
					}
					if !reflect.DeepEqual(request.body[key], want) {
						t.Fatalf("request JSON field %s changed: %v", key, request.body[key])
					}
				}
			case <-time.After(time.Second):
				t.Fatal("real SDK did not send the expected fixture request")
			}
			if len(requests) != 0 {
				t.Fatal("client unexpectedly repeated the fixture request")
			}
		})
	}
}

func TestHTTPErrorKeepsThePublicErrorContract(t *testing.T) {
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 429, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"error":{"type":"synthetic_type","message":"synthetic failure"}}`)), Request: r}, nil
	})
	c := NewClient("synthetic-key", upstream.WithHTTPClient(&http.Client{Transport: transport}))
	_, err := c.Engines(context.Background())
	var apiError APIError
	if !errors.As(err, &apiError) || apiError.StatusCode != 429 || apiError.Type != "synthetic_type" || apiError.Message != "synthetic failure" || err.Error() != "[429:synthetic_type] synthetic failure" {
		t.Fatalf("public APIError contract changed: %v", err)
	}
}

func TestTransportErrorIsNotReplaced(t *testing.T) {
	fault := errors.New("synthetic transport failure")
	transport := roundTripFunc(func(*http.Request) (*http.Response, error) { return nil, fault })
	c := NewClient("synthetic-key", upstream.WithHTTPClient(&http.Client{Transport: transport}))
	_, err := c.Engines(context.Background())
	if !errors.Is(err, fault) {
		t.Fatalf("transport error identity was lost: %v", err)
	}
}

func TestRequestCancellationReachesTheActualTransport(t *testing.T) {
	started := make(chan struct{})
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		close(started)
		<-r.Context().Done()
		return nil, r.Context().Err()
	})
	c := NewClient("synthetic-key", upstream.WithHTTPClient(&http.Client{Transport: transport, Timeout: 2 * time.Second}))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() { _, err := c.Engines(ctx); done <- err }()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("request did not reach transport")
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("request cancellation changed: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("cancelled request did not return")
	}
}

func TestStreamCancellationStopsTheRealSDKDecoder(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	writerDone := make(chan struct{})
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		reader, writer := io.Pipe()
		go func() {
			defer close(writerDone)
			if _, err := io.WriteString(writer, "data: {\"choices\":[{\"text\":\"synthetic piece\"}]}\n\n"); err != nil {
				_ = writer.CloseWithError(err)
				return
			}
			select {
			case <-r.Context().Done():
				_ = writer.CloseWithError(r.Context().Err())
			case <-time.After(2 * time.Second):
				_ = writer.CloseWithError(errors.New("synthetic stream cancellation timed out"))
			}
		}()
		return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"text/event-stream"}}, Body: reader, Request: r}, nil
	})
	c := NewClient("synthetic-key", upstream.WithBaseURL("https://fixture.invalid"), upstream.WithHTTPClient(&http.Client{Transport: transport, Timeout: 3 * time.Second}))
	pieces := ""
	err := c.CompletionStream(ctx, CompletionRequest{Prompt: []string{"synthetic prompt"}}, func(result *CompletionResponse) {
		for _, choice := range result.Choices {
			pieces += choice.Text
		}
		cancel()
	})
	if pieces != "synthetic piece" || !errors.Is(err, context.Canceled) {
		t.Fatalf("stream callback/cancellation contract changed: %q, %v", pieces, err)
	}
	select {
	case <-writerDone:
	case <-time.After(time.Second):
		t.Fatal("synthetic stream writer did not stop after cancellation")
	}
}
