package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/urfave/negroni"
)

var (
	categoryMap = map[string]int{
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

	client = &http.Client{
		Timeout: 10 * time.Second,
	}
)

type APIResponse struct {
	ResponseCode int        `json:"response_code"`
	Results      []Question `json:"results"`
}

type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
	Points           int
}

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
	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the welcome page!\n")
	})

	mux.HandleFunc("/home", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the about page!\n")
	})

	mux.HandleFunc("/trivia", triviaHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe(":8080", n))
}
