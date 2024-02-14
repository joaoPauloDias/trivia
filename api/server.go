// Package main provides the main server implementation for a trivia API.
// It fetches trivia questions from an external API and serves them over HTTP.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/urfave/negroni"
)

// categoryMap maps category names to their corresponding IDs used in the external API.
var categoryMap = map[string]int{
	"General Knowledge":                     9,
	"Entertainment: Books":                  10,
	"Entertainment: Film":                   11,
	"Entertainment: Music":                  12,
	"Entertainment: Musicals & Theatres":    13,
	"Entertainment: Television":             14,
	"Entertainment: Video Games":            15,
	"Entertainment: Board Games":            16,
	"Science & Nature":                      17,
	"Science: Computers":                    18,
	"Science: Mathematics":                  19,
	"Mythology":                             20,
	"Sports":                                21,
	"Geography":                             22,
	"History":                               23,
	"Politics":                              24,
	"Art":                                   25,
	"Celebrities":                           26,
	"Animals":                               27,
	"Vehicles":                              28,
	"Entertainment: Comics":                 29,
	"Science: Gadgets":                      30,
	"Entertainment: Japanese Anime & Manga": 31,
	"Entertainment: Cartoon & Animations":   32,
}

// client is the HTTP client used to make API requests.
var client = &http.Client{
	Timeout: 10 * time.Second,
}

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

// mapDifficultyToPoints maps the difficulty level of each question to a corresponding point value.
func mapDifficultyToPoints(questions []Question) {
	for i, question := range questions {
		switch question.Difficulty {
		case "easy":
			questions[i].Points = 1
		case "medium":
			questions[i].Points = 2
		case "hard":
			questions[i].Points = 3
		default:
			questions[i].Points = 0
		}
	}
}

// fetchTriviaQuestions fetches trivia questions from the external API.
// The amount parameter specifies the number of questions to fetch.
// The categoryId parameter specifies the category ID of the questions to fetch.
// If categoryId is empty, questions from all categories will be fetched.
// Returns the fetched questions and any error encountered.
func fetchTriviaQuestions(amount string, categoryId string) ([]Question, error) {
	params := url.Values{}
	params.Add("amount", amount)

	if categoryId != "" {
		params.Add("category", categoryId)
	}
	apiUrl := "https://opentdb.com/api.php?" + params.Encode()

	resp, err := client.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call returned non-200 status: %d", resp.StatusCode)
	}

	var apiResponse APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	mapDifficultyToPoints(apiResponse.Results)

	return apiResponse.Results, nil
}

// triviaHandler handles the HTTP request for trivia questions.
// It expects the "amount" and "category" query parameters.
// If "amount" is not provided, it defaults to 3.
// If "category" is not provided, questions from all categories will be fetched.
// Returns the fetched questions as a JSON response.
func triviaHandler(w http.ResponseWriter, req *http.Request) {
	amount := req.URL.Query().Get("amount")
	if amount == "" {
		amount = "3"
	}
	categoryName := req.URL.Query().Get("category")
	categoryId := ""
	if id, ok := categoryMap[categoryName]; ok {
		categoryId = fmt.Sprintf("%d", id)
	}

	questions, err := fetchTriviaQuestions(amount, categoryId)
	if err != nil {
		http.Error(w, "Failed to fetch trivia questions", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func main() {
	port := os.Getenv("API_PORT")
	if(port == ""){
		port = "3000"
	}
	port = fmt.Sprintf(":%s", port)
	fmt.Println("Server running on port", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/trivia", triviaHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe(port, n))
}
