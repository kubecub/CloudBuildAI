package client

import (
	"context"

)



// A Client is an API client to communicate with the OpenAI gpt-3 APIs
type Client interface {
	// Engines lists the currently available engines, and provides basic information about each
	// option such as the owner and availability.
	Engines(ctx context.Context) (*EnginesResponse, error)

	// Engine retrieves an engine instance, providing basic information about the engine such
	// as the owner and availability.
	Engine(ctx context.Context, engine string) (*EngineObject, error)

	// ChatCompletion creates a completion with the Chat completion endpoint which
	// is what powers the ChatGPT experience.
	ChatCompletion(ctx context.Context, request ChatCompletionRequest) (*ChatCompletionResponse, error)

	// ChatCompletion creates a completion with the Chat completion endpoint which
	// is what powers the ChatGPT experience.
	ChatCompletionStream(ctx context.Context, request ChatCompletionRequest, onData func(*ChatCompletionStreamResponse)) error

	// Completion creates a completion with the default engine. This is the main endpoint of the API
	// which auto-completes based on the given prompt.
	Completion(ctx context.Context, request CompletionRequest) (*CompletionResponse, error)

	// CompletionStream creates a completion with the default engine and streams the results through
	// multiple calls to onData.
	CompletionStream(ctx context.Context, request CompletionRequest, onData func(*CompletionResponse)) error

	// CompletionWithEngine is the same as Completion except allows overriding the default engine on the client
	CompletionWithEngine(ctx context.Context, engine string, request CompletionRequest) (*CompletionResponse, error)

	// CompletionStreamWithEngine is the same as CompletionStream except allows overriding the default engine on the client
	CompletionStreamWithEngine(ctx context.Context, engine string, request CompletionRequest, onData func(*CompletionResponse)) error

	// Given a prompt and an instruction, the model will return an edited version of the prompt.
	Edits(ctx context.Context, request EditsRequest) (*EditsResponse, error)

	// Search performs a semantic search over a list of documents with the default engine.
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// SearchWithEngine performs a semantic search over a list of documents with the specified engine.
	SearchWithEngine(ctx context.Context, engine string, request SearchRequest) (*SearchResponse, error)

	// Returns an embedding using the provided request.
	Embeddings(ctx context.Context, request EmbeddingsRequest) (*EmbeddingsResponse, error)

	// Moderation performs a moderation check on the given text against an OpenAI classifier to determine whether the
	// provided content complies with OpenAI's usage policies.
	Moderation(ctx context.Context, request ModerationRequest) (*ModerationResponse, error)
}
