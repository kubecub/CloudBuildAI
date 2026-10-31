// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package client

import upstream "github.com/PullRequestInc/go-gpt3"

// Preserve the existing data shapes through the source-compatible published SDK.
type (
	APIError                           = upstream.APIError
	APIErrorResponse                   = upstream.APIErrorResponse
	EngineObject                       = upstream.EngineObject
	EnginesResponse                    = upstream.EnginesResponse
	ChatCompletionRequestMessage       = upstream.ChatCompletionRequestMessage
	ChatCompletionRequest              = upstream.ChatCompletionRequest
	CompletionRequest                  = upstream.CompletionRequest
	EditsRequest                       = upstream.EditsRequest
	EmbeddingsRequest                  = upstream.EmbeddingsRequest
	LogprobResult                      = upstream.LogprobResult
	ChatCompletionResponseMessage      = upstream.ChatCompletionResponseMessage
	ChatCompletionResponseChoice       = upstream.ChatCompletionResponseChoice
	ChatCompletionStreamResponseChoice = upstream.ChatCompletionStreamResponseChoice
	ChatCompletionsResponseUsage       = upstream.ChatCompletionsResponseUsage
	ChatCompletionResponse             = upstream.ChatCompletionResponse
	ChatCompletionStreamResponse       = upstream.ChatCompletionStreamResponse
	CompletionResponseChoice           = upstream.CompletionResponseChoice
	CompletionResponse                 = upstream.CompletionResponse
	CompletionResponseUsage            = upstream.CompletionResponseUsage
	EditsResponse                      = upstream.EditsResponse
	EmbeddingsResult                   = upstream.EmbeddingsResult
	EmbeddingsUsage                    = upstream.EmbeddingsUsage
	EmbeddingsResponse                 = upstream.EmbeddingsResponse
	EditsResponseChoice                = upstream.EditsResponseChoice
	EditsResponseUsage                 = upstream.EditsResponseUsage
	SearchRequest                      = upstream.SearchRequest
	SearchData                         = upstream.SearchData
	SearchResponse                     = upstream.SearchResponse
	ModerationRequest                  = upstream.ModerationRequest
	ModerationCategoryResult           = upstream.ModerationCategoryResult
	ModerationCategoryScores           = upstream.ModerationCategoryScores
	ModerationResult                   = upstream.ModerationResult
	ModerationResponse                 = upstream.ModerationResponse
)
