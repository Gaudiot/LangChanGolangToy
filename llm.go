package main

import "github.com/tmc/langchaingo/llms/openai"

//MARK: - LLM

func NewOpenAILLM(api_key string) (*openai.LLM, error) {
	llm, err := openai.New(
		openai.WithToken(api_key),
		openai.WithModel("gpt-4o-mini"),
	)
	if err != nil {
		return nil, err
	}

	return llm, nil
}
