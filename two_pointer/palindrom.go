package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	start := 0
	end := len(runes) - 1
	for start < end {
		if !unicode.IsLetter(runes[start]) && !unicode.IsDigit(runes[start]) {
			start++
			continue
		}
		if !unicode.IsLetter(runes[end]) && !unicode.IsDigit(runes[end]) {
			end--
			continue
		}
		if unicode.ToLower(runes[start]) != unicode.ToLower(runes[end]) {
			return false
		}
		start++
		end--
	}
	return true
}