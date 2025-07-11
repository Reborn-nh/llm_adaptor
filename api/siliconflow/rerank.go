// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package siliconflow

import (
	"github.com/Reborn-nh/llm_adaptor/common"
)

type (
	CreateRerankReq struct {
		Model     string   `json:"model"`
		Query     string   `json:"query" `
		Documents []string `json:"documents"`
		TopK      int      `json:"top_k"`
	}
	CreateRerankRes struct {
		Results []*Data    `json:"results"`
		Message string     `json:"message"`
		Code    int        `json:"code"`
		Meta    RerankMeta `json:"meta"`
	}
	RerankMeta struct {
		BilledUnits struct {
			InputTokens     int `json:"input_tokens"`
			OutputTokens    int `json:"output_tokens"`
			SearchUnits     int `json:"search_units"`
			Classifications int `json:"classifications"`
		} `json:"billed_units"`
		Tokens struct {
			InputTokens  int `json:"input_tokens"`
			OutputTokens int `json:"output_tokens"`
		} `json:"tokens"`
	}
	Data struct {
		Index          int     `json:"index"`
		RelevanceScore float64 `json:"relevance_score"`
	}
)

func (r *CreateRerankReq) validate() error {
	return nil
}
func (c *Client) CreateRerank(req *CreateRerankReq) (*CreateRerankRes, error) {
	result := &CreateRerankRes{}
	if err := req.validate(); err != nil {
		return result, err
	}
	url := c.EndPoint + "/v1/rerank"
	headers := []common.Header{
		{Key: "Authorization", Value: "Bearer " + c.APIKey},
		{Key: "Context-Type", Value: "application/json"},
	}
	responseRaw, err := common.HttpPost(url, headers, nil, req)
	if err != nil {
		return result, err
	}
	err = common.HttpCheckError(responseRaw, &ErrorResponse{})
	if err != nil {
		return result, err
	}

	err = common.HttpDecodeResponse(responseRaw, &result)
	if err != nil {
		return result, err
	}
	return result, err
}
