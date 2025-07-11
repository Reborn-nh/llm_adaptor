package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/Reborn-nh/llm_adaptor/adaptor"
)

func testChatCompletion(Meta adaptor.Meta) {
	client := &adaptor.Adaptor{}
	client.Init(Meta)
	req := adaptor.ZhimaChatCompletionRequest{
		Messages:    []adaptor.ZhimaChatCompletionMessage{{Role: "user", Content: "你好"}},
		Temperature: 0.1,
		MaxToken:    10,
	}
	res, err := client.CreateChatCompletion(req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res.Result)
}

func TestOpenAIChatCompletion(t *testing.T) {
	testChatCompletion(adaptor.Meta{
		Corp:   "openai",
		Model:  `gpt-3.5-turbo`,
		APIKey: os.Getenv(`OPENAI_KEY`),
	})
}

func TestMinimaxiChatCompletion(t *testing.T) {
	testChatCompletion(adaptor.Meta{
		Corp:   "minimax",
		Model:  `abab6.5s-chat`,
		APIKey: os.Getenv(`MINIMAX_KEY`),
	})
}
func TestSiliconFlowChatCompletion(t *testing.T) {
	testChatCompletion(adaptor.Meta{
		Corp:       "siliconflow",
		EndPoint:   `https://api.siliconflow.cn`,
		APIVersion: "v1",
		Model:      `Qwen/Qwen2.5-72B-Instruct`,
		APIKey:     os.Getenv(`SILICONFLOW_KEY`),
	})
}
