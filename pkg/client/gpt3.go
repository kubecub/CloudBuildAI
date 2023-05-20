// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package client

import (
	"net/http"
	"time"
)

type client struct {
	baseURL       string
	apiKey        string
	userAgent     string
	httpClient    *http.Client
	defaultEngine string
	idOrg         string
}

// NewClient returns a new OpenAI GPT-3 API client. An apiKey is required to use the client
func NewClient(apiKey string, options ...ClientOption) Client {
	httpClient := &http.Client{
		Timeout: time.Duration(defaultTimeoutSeconds * time.Second),
	}

	c := &client{
		userAgent:     defaultUserAgent,
		apiKey:        apiKey,
		baseURL:       defaultBaseURL,
		httpClient:    httpClient,
		defaultEngine: DefaultEngine,
		idOrg:         "",
	}
	for _, o := range options {
		o(c)
	}
	return c
}
