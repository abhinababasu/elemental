package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abhinababasu/elemental/element"
)

type ElementResponse struct {
	Word     string
	Elements []element.Element
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	w.Header().Add("Access-Control-Allow-Origin", "*")
	words := q.Get("words")
	if len(words) > 200 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Query too long"))
		return
	}
	if len(words) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Expected query parameters ?words=<words>"))
		log.Println("No words specified")
		return
	}

	i := 0
	result := make([]ElementResponse, 0)
	for _, word := range strings.Split(words, " ") {
		i++
		log.Println("Word: ", i, word)
		if len(word) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Too long words"))
			return
		}

		res := element.GetElements().GetElementsForWord(word)
		elementRes := ElementResponse{Word: word, Elements: res}
		result = append(result, elementRes)
		log.Println("Found elements", len(res))
	}

	j, _ := json.MarshalIndent(result, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	port = ":" + port

	http.HandleFunc("/", handler)
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
