package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

//MARK: - CUSTOM TOOL

type SubjectLevel string

const (
	beginner     SubjectLevel = "beginner"
	intermediate SubjectLevel = "intermediate"
	advanced     SubjectLevel = "advanced"
)

type UserSubjectLevel map[string]SubjectLevel

type UserLevelTool struct {
	levels map[string]UserSubjectLevel
}

type UserLevelInput struct {
	UserID  string `json:"user_id"`
	Subject string `json:"subject"`
}

func NewUserLevelTool() *UserLevelTool {
	return &UserLevelTool{
		levels: map[string]UserSubjectLevel{
			"id1": {
				"history": beginner,
				"math":    advanced,
				"science": advanced,
				"english": intermediate,
			},
			"id2": {
				"history":   advanced,
				"spanish":   advanced,
				"math":      intermediate,
				"arts":      beginner,
				"chemistry": beginner,
			},
			"id3": {
				"arts":      advanced,
				"chemistry": intermediate,
			},
		},
	}
}

func (ult *UserLevelTool) Name() string {
	return "get_user_level_tool"
}

func (ult *UserLevelTool) Description() string {
	return `Returns the learning level of a user. 
	Possible values are beginner, intermediate, or advanced. 
	Input must be valid JSON with this structure:

	{
		"user_id": "string",
		"subject": "string"
	}`
}

func (ult *UserLevelTool) Call(ctx context.Context, input string) (string, error) {
	var params UserLevelInput

	if err := json.Unmarshal([]byte(input), &params); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	userId, subject := params.UserID, params.Subject
	subject = strings.ToLower(subject)
	subject = strings.TrimSpace(subject)

	userLevels, ok := ult.levels[userId]
	if !ok {
		return "", fmt.Errorf("user not found: %s", userId)
	}

	level, ok := userLevels[subject]
	if !ok {
		return string(beginner), nil // defaults to beginner
	}

	return string(level), nil
}
