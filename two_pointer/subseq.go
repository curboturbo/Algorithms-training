package main

import "unicode"

func isSubsequence(s string, t string) bool {
    if s == ""{return true}
	a := []rune(s)
	b := []rune(t)
	var pointer int = 0
	for i := 0; i < len(b); i++ {
		if unicode.ToLower(b[i]) == unicode.ToLower(a[pointer]) {
			pointer++
		}
		if pointer == len(a) {
			return true
		}
	}
	return false
}

