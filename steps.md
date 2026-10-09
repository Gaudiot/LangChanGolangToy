# Step 0 - Set Up the Go Project

Before getting started with LangChain, we need to initialize a Go project and set up its basic structure.

### Initialize the Project

Create a Go module by running:

```bash
go mod init github.com/{username}/langchain
```

### Create the Main File

Create a `main.go` file in the project root. This will serve as the entry point for our application.

### Configure Environment Variables

Create a `.env` file to store the API keys and other configuration values required by the application.

Make sure to add `.env` to your `.gitignore` to avoid accidentally committing sensitive information.

# Step 1 - Install Dependencies

Next, we'll install the packages needed to build our LangChain application, integrate external tools, and manage configuration.

### LangChain for Go

The main library we'll use throughout this project.

```bash
go get github.com/tmc/langchaingo
```

### OpenAI Integration

Provides the integration between LangChain and OpenAI's language models.

```bash
go get github.com/tmc/langchaingo/llms/openai
```

### Prompt Templates

Allows us to define reusable prompts with dynamic variables.

```bash
go get github.com/tmc/langchaingo/prompts
```

### LangChain Callbacks

Provides callbacks that can be used to observe different stages of the chain's execution.

```bash
go get github.com/tmc/langchaingo/callbacks
```

### Model Context Protocol (MCP)

The SDK we'll use to connect our application to external tools through MCP.

```bash
go get github.com/modelcontextprotocol/go-sdk/mcp
```

### Environment Variable Loader

Loads environment variables from our `.env` file.

```bash
go get github.com/joho/godotenv
```

# Step 2 - Create a Custom Tool

One of the key features of LangChain agents is their ability to use external tools to retrieve information or perform specific actions.

For this project, we'll create a custom tool that retrieves a user's knowledge level for a given subject.

The tool receives a **user ID** and a **subject**, and returns one of three possible knowledge levels: `beginner`, `intermediate`, or `advanced`.

We'll also define some basic rules for its behavior:

- If the user doesn't exist, the tool returns an error.
- If the user exists but has no knowledge level registered for the requested subject, the tool defaults to `beginner`.

Since our goal is to understand how custom tools work, we'll use mocked data instead of connecting to a real database or external service.

# Step 3 - Define the Expected Response

Before building our chain, we'll define the Go struct that represents the expected output of our flashcard generator.

This gives us a clear structure for the data we want to receive from the language model.

We'll use this struct later when implementing an output parser to convert the model's response into a Go value that our application can work with.

# Step 4 - Create a Prompt Template

Now we need to define the instructions that will guide the language model when generating flashcards.

Instead of writing a fixed prompt, we'll use a **prompt template** that supports dynamic variables. This allows us to reuse the same instructions while adjusting values such as the subject, the user's knowledge level, and other relevant information.

### Escaping Curly Brackets

When including JSON examples or other literal structures containing curly brackets in a prompt template, make sure to escape them using double curly brackets: `{{` and `}}`.

Otherwise, LangChain may interpret the content inside the brackets as template variables.

### Defining Message Roles

We'll also apply some basic prompt engineering techniques to make our instructions clearer and more consistent.

In particular, we'll use two message roles:

- **System:** Defines the model's general behavior, responsibilities, and constraints.
- **Human:** Represents the user's request and provides the context for generating flashcards.

Other message roles are available, but these two are enough for our introductory project.

# Step 5 - Initialize the Language Model

With our prompt template ready, we need to configure the language model that will process our requests.

We'll use LangChain's OpenAI integration to initialize an LLM and define the necessary configuration, such as the model and its API credentials.

This model will be responsible for interpreting our prompts and generating the flashcard content.

# Step 6 - Create the Agent and Executor

Now we'll connect the language model, prompt, and available tools using an **agent** and an **executor**.

These components have different responsibilities:

- **Agent:** Uses the language model to decide how to approach a request, including whether an available tool should be called.
- **Executor:** Manages the execution process, invokes the tools selected by the agent, and returns their results to the agent when needed.

Together, they allow our application to do more than simply send a prompt to an LLM. The agent can make decisions and interact with tools before producing its final response.

# Step 7 - Integrate an MCP Tool

So far, our agent can use the custom tool we created to retrieve a user's knowledge level.

Now we'll extend its capabilities by integrating an external tool through the **Model Context Protocol (MCP)**.

MCP provides a standardized way for AI applications to connect with external tools and data sources.

For this project, we'll connect to a Wikipedia MCP server, allowing our agent to retrieve additional information about a subject.

This gives the agent access to an external source of factual information that can help improve the quality and accuracy of the generated flashcards.

# Step 8 - Add Conversation Memory

Another useful feature of LangChain is the ability to maintain context across multiple interactions.

By adding **conversation memory** to our agent executor, we can preserve information from previous exchanges and make it available in subsequent requests.

It's important to understand that, in our implementation, this memory belongs to a specific executor instance. It isn't automatically shared between different executor instances.

In this toy project, memory doesn't provide much practical benefit because each execution consists of a single user interaction.

However, implementing it is still a useful exercise to understand how LangChain can support conversations that depend on previous context.

# Step 9 - Implement an Output Parser

Language models typically return their responses as text, but applications often need to work with structured data.

To address this, we'll introduce an **output parser** that converts the model's response into a more useful format.

In our case, we'll use the struct defined in Step 3 to represent the generated flashcards as structured Go data.

This makes it easier to access individual fields and use the generated content elsewhere in the application.

# Step 10 - Add Observability

Now that our application is working, we'll add **observability** to better understand what happens during execution.

LangChain provides callback mechanisms that allow us to inspect different stages of the process, such as model interactions, tool calls, and execution results.

This is particularly useful when working with agents because their execution flow isn't always predictable.

By observing these interactions, we can better understand the agent's decisions, identify unexpected behavior, and debug potential issues.

# Bonus - Build a Simple Chain Without an Agent

Not every LangChain application needs an agent.

We can also create a **simple chain** that follows a predefined sequence of operations.

Unlike an agent, which can dynamically decide whether to call tools or take additional steps, a simple chain follows a fixed execution flow from input to output.

This is a useful way to understand the distinction between deterministic workflows and agent-based applications, as well as when each approach might be appropriate.