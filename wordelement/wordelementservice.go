package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/abhinababasu/elemental/element"
)

type ElementResponse struct {
	Word     string
	Elements []element.Element
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	log.Printf("Remote=%v;Request=%v %v,UA=%v", r.RemoteAddr, r.Method, r.URL.String(), r.UserAgent())

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

	timeEnd := time.Since(timeStart)
	log.Printf("Solved=%v;time=%vms;words=%v", words, timeEnd.Milliseconds(), len(result.Result))
}

func main() {
	port := flag.Int("port", 8080, "Port server will listen on")
	flag.Parse()

	http.HandleFunc("/", handler)
	addr := fmt.Sprintf(":%v", *port)
	log.Println("Listening on ", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
