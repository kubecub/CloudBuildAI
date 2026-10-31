// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package client

import upstream "github.com/PullRequestInc/go-gpt3"

type ClientOption = upstream.ClientOption

// NewClient returns the original SDK client through the existing public interface.
func NewClient(apiKey string, options ...ClientOption) Client {
	return upstream.NewClient(apiKey, options...)
}

// IntPtr retains the helper used by the existing chat example.
func IntPtr(value int) *int {
	return upstream.IntPtr(value)
}
