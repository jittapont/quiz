package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) == word
}

func main() {
	fmt.Println(isPalindrome("Deleveled"))
}
