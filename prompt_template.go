package main

import "github.com/tmc/langchaingo/prompts"

// MARK: - PROMPT TEMPLATE

func NewFlascardPrompt() prompts.PromptTemplate {
	return prompts.PromptTemplate{
		Template: `
		System:
		You are an educational assistant.
		Always determine the user's knowledge level before generating flashcards.
		Use Wikipedia when additional factual information would improve accuracy.
		Adapt the difficulty to the user's level.

		Return exactly 10 flashcards as valid JSON.
		Return ONLY valid JSON using exactly this structure:
		{{
			"subject": "string",
			"user_level": "beginner | intermediate | advanced",
			"flashcards": [
				{{
				"question": "string",
				"answer": "string"
				}}
			]
		}}	
		Do not include markdown fences or any text outside the JSON.

		Human:
		Create flashcards about {subject}.

		The current user ID is {user_id}.
		`,
		InputVariables: []string{
			"subject",
			"user_id",
		},
		TemplateFormat: prompts.TemplateFormatFString,
	}
}
