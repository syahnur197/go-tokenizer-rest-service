package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jdkato/prose/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/top-ten-words/", handleTopTenWords)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func handleTopTenWords(w http.ResponseWriter, r *http.Request) {
	setJson(w)

	text := r.FormValue("text")

	doc, err := prose.NewDocument(text)

	if err != nil {
		responseForbidden(w, "Fail to tokenize")
	}

	tokensCount := countTokens(doc.Tokens())

	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Test",
		Data:    tokensCount,
	})
}

func responseSuccess(w http.ResponseWriter, message string) {
	response := Response{
		Success: true,
		Message: message,
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func responseForbidden(w http.ResponseWriter, message string) {
	response := Response{
		Success: false,
		Message: message,
	}
	w.WriteHeader(http.StatusForbidden)

	json.NewEncoder(w).Encode(response)
}

func setJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func countTokens(tok []prose.Token) map[string]int {
	tokenCounts := make(map[string]int)

	for _, token := range tok {
		tokenText := strings.ToLower(token.Text)

		if tokenIsNotWord(token) {
			continue
		}

		tokenCounts = addTokenCount(tokenCounts, tokenText)
	}

	return tokenCounts

}

func tokenIsNotWord(token prose.Token) bool {
	tokenTag := token.Tag

	isNotWord := tokenTag == "(" ||
		tokenTag == ")" ||
		tokenTag == "," ||
		tokenTag == ":" ||
		tokenTag == "." ||
		tokenTag == "''" ||
		tokenTag == "``" ||
		tokenTag == "#" ||
		tokenTag == "$"

	return isNotWord
}

func addTokenCount(tokenCounts map[string]int, tokenText string) map[string]int {
	if val, ok := tokenCounts[tokenText]; ok {
		tokenCounts[tokenText] = val + 1
	} else {
		tokenCounts[tokenText] = 1
	}

	return tokenCounts
}
