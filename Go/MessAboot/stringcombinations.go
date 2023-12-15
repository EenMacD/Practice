package main

import (
	"fmt"
	"strings"
)

func generateWordCombinations(letters string, prefix string, results *[]string) {
	if len(letters) == 0 {
		*results = append(*results, prefix)
		return
	}

	for i := 0; i < len(letters); i++ {
		nextLetter := string(letters[i])
		remainingLetters := letters[:i] + letters[i+1:]
		generateWordCombinations(remainingLetters, prefix+nextLetter, results)
	}
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	input = strings.ToLower(input) // Convert to lowercase for case-insensitivity
	var results []string

	generateWordCombinations(input, "", &results)

	fmt.Println("Word combinations:")
	for _, word := range results {
		fmt.Println(word)
	}
}
