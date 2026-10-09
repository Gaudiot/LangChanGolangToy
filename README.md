# Flashcard Generator with LangChain and Go

A toy project built to explore the fundamentals of **LangChain using Go (Golang)**.

The idea is to build a simple AI-powered chain that generates flashcards to help users study a subject of their choice. Along the way, the project explores some of LangChain's core concepts, including chains, tool integration, memory, and observability.

This is not intended to be a production-ready application or a showcase of software engineering best practices. It's simply a hands-on learning project focused on understanding how LangChain works and how to implement its basic features in Go.

## Features

### AI-Powered Flashcard Generation

The main goal of the application is to generate study flashcards based on a subject provided by the user.

Rather than relying exclusively on the language model's existing knowledge, the chain can use external tools to gather additional context and tailor the generated content.

### Custom Tool: User Knowledge Level

A custom tool retrieves a user's knowledge level for a specific subject.

Given a **user ID** and a **subject**, the tool provides information about the user's familiarity with that topic.

This allows the chain to adjust the complexity of the generated flashcards based on the user's knowledge level.

### Wikipedia Integration via MCP

The project integrates a Wikipedia tool through the **Model Context Protocol (MCP)**.

This allows the chain to retrieve additional information about a subject, helping it generate flashcards grounded in factual information rather than relying solely on the model's internal knowledge.

### Conversation Memory

The chain includes memory to retain context from previous interactions.

This makes it possible to explore how conversation history can influence subsequent responses and how LangChain manages context across multiple interactions.

### Observability

The project also includes observability capabilities to help inspect and understand the chain's execution.

This makes it easier to follow the flow of interactions, including model calls, tool usage, and the overall execution process.

## Learning Goals

Through this project, I wanted to get hands-on experience with:

- Building a basic LangChain application in Go.
- Understanding how chains orchestrate interactions with language models.
- Creating and integrating custom tools.
- Connecting external tools through MCP.
- Using conversation memory to maintain context.
- Exploring observability to understand what's happening during chain execution.

## How It Works

At a high level, the application follows this flow:

1. The user provides a subject they want to study.
2. The chain can retrieve the user's knowledge level through a custom tool.
3. The chain can query Wikipedia through MCP to gather additional factual information.
4. The language model uses the available context to generate flashcards tailored to the subject and the user's knowledge level.
5. Conversation memory helps preserve relevant context across interactions.
6. Observability allows the execution flow to be inspected.

The goal is not to build a complex flashcard application, but to use this scenario as a practical way to understand how LangChain's different components work together.

## Getting Started

The project was developed incrementally, with each step introducing a new concept or feature.

For a walkthrough of the implementation and the learning process, check out the [step-by-step guide](./steps.md).

## Disclaimer

This is an **introductory learning project**.

It does not aim to follow every software engineering best practice, cover all of LangChain's capabilities, or provide a production-ready architecture.

The focus is on experimentation, learning by doing, and gaining a practical understanding of how to work with LangChain in Go.