package test

import "strings"

func Greeting() string {
	// exported
	return "Greetings"
}
func WordCount(s string) map[string]int {
	// exported
	wordCountMap := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		elem, ok := wordCountMap[word]
		if ok {
			wordCountMap[word] = elem + 1
		} else {
			wordCountMap[word] = 1
		}
	}
	return wordCountMap
}
