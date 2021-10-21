package github

import "encoding/json"

// Context holds all the information set on Github runners in relation to the job
//
// This information is retrieved from variables during workflow execution
type Context struct {
	Action          string          `json:"action"`
	ActionPath      string          `json:"action_path"`
	Actor           string          `json:"actor"`
	BaseRef         string          `json:"base_ref"`
	Event           json.RawMessage `json:"event"`
	EventName       string          `json:"event_name"`
	EventPath       string          `json:"event_path"`
	HeadRef         string          `json:"head_ref"`
	Job             string          `json:"job"`
	Ref             string          `json:"ref"`
	Repository      string          `json:"repository"`
	RepositoryOwner string          `json:"repository_owner"`
	RunID           string          `json:"run_id"`
	RunNumber       string          `json:"run_number"`
	SHA             string          `json:"sha"`
	Token           Token           `json:"token,omitempty"`
	Workflow        string          `json:"workflow"`
	Workspace       string          `json:"workspace"`
}

// Token the github token used during a workflow
type Token string

// UnmarshalText Unmarshals the token received from Github
func (t *Token) UnmarshalText(text []byte) error {
	*t = Token(text)
	return nil
}

// MarshalText masks the token as *** when marshalling
func (t Token) MarshalText() ([]byte, error) {
	return []byte("***"), nil
}

// RunnerContext holds information about the given Github Runner in which a workflow executes
//
// This information is retrieved from variables during workflow execution
type RunnerContext struct {
	OS        string `json:"os"`
	Temp      string `json:"temp"`
	ToolCache string `json:"tool_cache"`
}

// AnyEvent holds the inputs from a Github workflow
//
// See https://docs.github.com/en/actions/reference/events-that-trigger-workflows
// The only Event with dynamically-provided input is workflow_dispatch which
// exposes the user params at the key "input."
type AnyEvent struct {
	Inputs json.RawMessage `json:"inputs"`
}
