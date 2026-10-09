package main

import (
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

// MARK: - Agent

func NewAgent(llm *openai.LLM, tools []tools.Tool) *agents.OpenAIFunctionsAgent {
	return agents.NewOpenAIFunctionsAgent(
		llm,
		tools,
	)
}
