// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package minimax

import "github.com/Reborn-nh/llm_adaptor/api/openai"

type Client struct {
	APIKey       string
	EndPoint     string
	OpenAIClient *openai.Client
}

func NewClient(APIKey string) *Client {
	return &Client{
		APIKey:   APIKey,
		EndPoint: "https://api.minimax.chat/v1",
		OpenAIClient: &openai.Client{
			EndPoint: "https://api.minimax.chat/v1",
			APIKey:   APIKey,
			ErrResp:  &openai.ErrorResponse{},
		},
	}
}
