package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// MARK: - MCP

func CreateWikipediaSession(
	ctx context.Context,
) (*mcp.ClientSession, error) {
	client := mcp.NewClient(
		&mcp.Implementation{
			Name:    "flashcard-agent",
			Version: "1.0.0",
		},
		nil,
	)

	transport := &mcp.StreamableClientTransport{
		Endpoint: "https://wikipedia.caseyjhand.com/mcp",
	}

	session, err := client.Connect(
		ctx,
		transport,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}

type WikipediaTool struct {
	session *mcp.ClientSession
}

func NewWikipediaTool(
	session *mcp.ClientSession,
) *WikipediaTool {
	return &WikipediaTool{
		session: session,
	}
}

func (wt *WikipediaTool) Name() string {
	return "wikipedia_mcp"
}

func (wt *WikipediaTool) Description() string {
	return `
	Search Wikipedia for factual information about a subject.

	Use this tool when additional factual information would
	improve the accuracy of the flashcards.

	Input should be the topic to search.
	`
}

func (wt *WikipediaTool) Call(
	ctx context.Context,
	input string,
) (string, error) {
	result, err := wt.session.CallTool(
		ctx,
		&mcp.CallToolParams{
			Name: "wikipedia_search_articles",
			Arguments: map[string]any{
				"query": input,
			},
		},
	)
	if err != nil {
		return "", err
	}

	if result.IsError {
		return "", fmt.Errorf("Wikipedia MCP returned an error")
	}

	var builder strings.Builder
	for _, content := range result.Content {
		textContent, ok := content.(*mcp.TextContent)
		if !ok {
			continue
		}

		builder.WriteString(textContent.Text)
		builder.WriteString("\n")
	}

	return builder.String(), nil
}

// MARK: - MCP List Tools

func (wt *WikipediaTool) ListTools(
	ctx context.Context,
) error {
	toolsResult, err := wt.session.ListTools(
		ctx,
		nil,
	)
	if err != nil {
		return err
	}

	for _, tool := range toolsResult.Tools {
		fmt.Println(tool.Name)
		fmt.Println(tool.Description)
		fmt.Println()
	}

	return nil
}
