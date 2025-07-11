// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package openrouter

import "github.com/Reborn-nh/llm_adaptor/api/openai"

type Client struct {
	APIKey       string
	EndPoint     string
	OpenAIClient *openai.Client
}

func NewClient(EndPoint, APIKey string) *Client {
	return &Client{
		APIKey:   APIKey,
		EndPoint: EndPoint,
		OpenAIClient: &openai.Client{
			EndPoint: EndPoint,
			APIKey:   APIKey,
			ErrResp:  &openai.ErrorResponse{},
		},
	}
}
