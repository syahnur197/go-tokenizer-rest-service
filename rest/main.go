package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jdkato/prose/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Token struct {
	Text  string `json:"text"`
	Count int    `json:"count"`
}

// Sortable Type
type TokensList []Token

func (list TokensList) Len() int           { return len(list) }
func (list TokensList) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }
func (list TokensList) Less(i, j int) bool { return list[i].Count < list[j].Count }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/top-ten-words/", handleTopTenWords).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func handleTopTenWords(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("text")

	doc, err := prose.NewDocument(text)

	if err != nil {
		responseJson(w, Response{
			Success: false,
			Message: "Fail to tokenize",
			Data:    nil,
		})

		return
	}

	tokensCount := countTokens(doc.Tokens())

	sortedTokens := sortTokensDesc(tokensCount)

	topTenSortedTokens := sortedTokens[:10]

	responseJson(w, Response{
		Success: true,
		Message: "Sorted successfully",
		Data:    topTenSortedTokens,
	})
}

func responseJson(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func countTokens(tokens []prose.Token) map[string]int {
	tokenCounts := make(map[string]int)

	// counting the tokens
	for _, token := range tokens {

		// lowercase all the tokens
		tokenText := strings.ToLower(token.Text)

		if tokenIsNotWord(token) {
			continue
		}

		tokenCounts = addTokenCount(tokenCounts, tokenText)
	}

	return tokenCounts

}

func sortTokensDesc(tokenCounts map[string]int) TokensList {
	tokensList := make(TokensList, len(tokenCounts))

	i := 0
	for key, value := range tokenCounts {
		tokensList[i] = Token{
			Text:  key,
			Count: value,
		}
		i++
	}

	sort.Sort(sort.Reverse(tokensList))

	return tokensList
}

func tokenIsNotWord(token prose.Token) bool {
	// some tokens are not actual words
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
	// add 1 to count if tokenText already exist as key in tokenCounts
	// otherwise, add new key tokenText with the value of 1
	if val, ok := tokenCounts[tokenText]; ok {
		tokenCounts[tokenText] = val + 1
	} else {
		tokenCounts[tokenText] = 1
	}

	return tokenCounts
}
