// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package openrouter

import "github.com/Reborn-nh/llm_adaptor/api/openai"

type Client struct {
	APIKey       string
	EndPoint     string
	OpenAIClient *openai.Client
}

func NewClient(APIKey string) *Client {
	return &Client{
		APIKey:   APIKey,
		EndPoint: "https://openrouter.ai/api/v1",
		OpenAIClient: &openai.Client{
			EndPoint: "https://openrouter.ai/api/v1",
			APIKey:   APIKey,
			ErrResp:  &openai.ErrorResponse{},
		},
	}
}
