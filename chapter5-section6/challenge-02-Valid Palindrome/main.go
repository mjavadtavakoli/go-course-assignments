package main

import (
	strg "strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strg.ToLower(s) // add alias

	left := 0
	right := len(s) - 1

	for left < right {

		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}

		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
