package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/schema"
)

// MARK: - OBSERVABILITY

type ObservabilityHandler struct {
	callbacks.SimpleHandler
}

func (h ObservabilityHandler) HandleChainStart(
	ctx context.Context,
	inputs map[string]any,
) {
	fmt.Println("[CHAIN START]")
	fmt.Println(inputs)
}

func (h ObservabilityHandler) HandleChainEnd(
	ctx context.Context,
	outputs map[string]any,
) {
	fmt.Println("[CHAIN END]")
	fmt.Println(outputs)
}

func (h ObservabilityHandler) HandleToolStart(
	ctx context.Context,
	input string,
) {
	fmt.Println("[TOOL START]")
	fmt.Println("input:", input)
}

func (h ObservabilityHandler) HandleToolEnd(
	ctx context.Context,
	output string,
) {
	fmt.Println("[TOOL END]")
	fmt.Println("output:", output)
}

func (h ObservabilityHandler) HandleToolError(
	ctx context.Context,
	err error,
) {
	fmt.Println("[TOOL ERROR]")
	fmt.Println(err)
}

func (h ObservabilityHandler) HandleAgentAction(
	ctx context.Context,
	action schema.AgentAction,
) {
	fmt.Println("[AGENT ACTION]")
	fmt.Println("tool:", action.Tool)
	fmt.Println("input:", action.ToolInput)
}

func (h ObservabilityHandler) HandleAgentFinish(
	ctx context.Context,
	finish schema.AgentFinish,
) {
	fmt.Println("[AGENT FINISH]")
	fmt.Println(finish.ReturnValues)
}

func (h ObservabilityHandler) HandleLLMStart(
	ctx context.Context,
	prompts []string,
) {
	// TODO
}

func (h ObservabilityHandler) HandleLLMError(
	ctx context.Context,
	err error,
) {
	// TODO
}
