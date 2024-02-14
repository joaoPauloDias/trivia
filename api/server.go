package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	port := getServerPort()
	logger := NewLogger()

	mux := http.NewServeMux()
	mux.HandleFunc("/trivia", func(w http.ResponseWriter, r *http.Request) {
		triviaHandler(w, r, logger)
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	logger.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(port, n))
}
