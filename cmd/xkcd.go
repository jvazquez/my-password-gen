package main

import (
	"fmt"
	"github.com/sethvargo/go-diceware/diceware"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func passwordGeneratorHandler(w http.ResponseWriter, r *http.Request) {
	var defaultGeneratedWords int = 6
	var seed int = 0

	words, ok := r.URL.Query()["words"]
	if !ok || len(words[0]) < 1 {
		log.Println("Url Param 'words' is missing.Using default for generated words",
			defaultGeneratedWords)
		seed = defaultGeneratedWords
	} else {
		seed, _ = strconv.Atoi(words[0])
	}
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(seed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, strings.Join(list, " "))
	fmt.Fprintf(w, "\n")
}

func main() {
	http.HandleFunc("/", passwordGeneratorHandler)
	http.ListenAndServe(":8000", nil)
}
