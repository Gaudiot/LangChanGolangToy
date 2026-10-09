package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
)

//MARK: - CHAIN RESPONSE MODEL

type Flashcard struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type FlashcardResponse struct {
	Subject    string      `json:"subject"`
	UserLevel  string      `json:"user_level"`
	Flashcards []Flashcard `json:"flashcards"`
}

// MARK: - OUTPUT PARSER

func OutputParser(input string) (*FlashcardResponse, error) {
	var response FlashcardResponse
	err := json.Unmarshal([]byte(input), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// MARK: - MAIN FUNCTION

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not defined")
	}

	llm, err := NewOpenAILLM(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	userLevelTool := NewUserLevelTool()
	wikipediaSession, err := CreateWikipediaSession(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if wikipediaSession == nil {
		log.Fatal("Wikipedia MCP session is nil")
	}
	defer wikipediaSession.Close()
	var wikipediaTool = NewWikipediaTool(wikipediaSession)
	agentTools := []tools.Tool{
		userLevelTool,
		wikipediaTool,
	}

	agent := NewAgent(
		llm,
		agentTools,
	)

	conversationMemory := memory.NewConversationBuffer()
	observer := ObservabilityHandler{}

	executor := agents.NewExecutor(
		agent,
		agents.WithMemory(conversationMemory),
		agents.WithCallbacksHandler(observer),
	)

	subject := "history"
	userID := "id3"
	promptTemplate := NewFlascardPrompt()
	input, err := promptTemplate.Format(
		map[string]any{
			"subject": subject,
			"user_id": userID,
		},
	)
	if err != nil {
		log.Fatal("input: ", err)
	}

	result, err := chains.Run(
		ctx,
		executor,
		input,
	)
	if err != nil {
		log.Fatal(err)
	}

	response, err := OutputParser(result)
	if err != nil {
		log.Fatalf("failed to parse agent response: %v\nresponse: %s\n", err, result)
	}

	fmt.Println(response.Subject)
	fmt.Println(response.UserLevel)
	fmt.Println(response.Flashcards[0].Answer + " - " + response.Flashcards[0].Question)
}

// MARK: - Simple Chain

func SimpleChain() {
	ctx := context.Background()

	// Step 1: Get env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not defined")
	}

	// Step 2: Get user level
	user_level, err := NewUserLevelTool().Call(ctx, "{'user_id': 'id1', 'subject': 'history'}")
	if err != nil {
		log.Fatal(err)
	}

	// Step 3: Start LLM
	llm, err := NewOpenAILLM(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Create Chain
	chain := chains.NewLLMChain(
		llm,
		NewFlascardPrompt(),
	)

	// Step 5: Call Chain
	result, err := chains.Call(
		ctx,
		chain,
		map[string]any{
			"subject":    "history",
			"user_level": user_level,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	output, ok := result["text"].(string)
	if !ok {
		log.Fatal("invalid chain output")
	}

	// Step 6: Parse Chain Output
	parser, err := OutputParser(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parser.Subject)
}
