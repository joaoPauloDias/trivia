package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var (
	// client is the HTTP client used to make API requests.
	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	// categoryMap maps category names to their corresponding IDs used in the external API.
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
)

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
