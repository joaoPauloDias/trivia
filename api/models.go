package main

// APIResponse represents the response from the trivia API.
type APIResponse struct {
	ResponseCode int        `json:"response_code"`
	Results      []Question `json:"results"`
}

// Question represents a trivia question.
type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
	Points           int      `json:"points"`
}
