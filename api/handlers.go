package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// triviaHandler handles the HTTP request for trivia questions.
// It expects the "amount" and "category" query parameters.
// If "amount" is not provided, it defaults to 3.
// If "category" is not provided, questions from random categories are fetched.
// Returns the fetched questions as a JSON response.
func triviaHandler(w http.ResponseWriter, req *http.Request, logger *log.Logger) {
	amount := req.URL.Query().Get("amount")
	if amount == "" {
		amount = defaultAmount
	}

	categoryName := req.URL.Query().Get("category")
	categoryId := ""
	if id, ok := categoryMap[categoryName]; ok {
		categoryId = fmt.Sprintf("%d", id)
	}

	startTime := time.Now()
	questions, err := fetchTriviaQuestions(amount, categoryId)
	elapsedTime := time.Since(startTime)

	if err != nil {
		http.Error(w, "Failed to fetch trivia questions", http.StatusInternalServerError)
		logger.Println(err)
		return
	}

	logger.Println("Time taken to fetch trivia questions from external API:", elapsedTime)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(questions)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		logger.Println(err)
		return
	}
}
