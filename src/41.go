package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)

	if len(words) < 0 {
		return result
	}

	for _, word := range words {
		result[word] += 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
