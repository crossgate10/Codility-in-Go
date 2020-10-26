package LongestPassword

import (
	"strings"
	"unicode"
)

func Solution(S string) int {
	words := strings.Split(S, " ")
	maxLen := -1
	for _, word := range words {
		if isValidPassword(word) {
			maxLen = max(maxLen, len(word))
		}
	}
	return maxLen
}

func isValidPassword(s string) bool {
	n := len(s)
	if n%2 == 0 {
		return false
	}
	digitCnt := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			continue
		}
		if unicode.IsDigit(c) {
			digitCnt++
			continue
		}
		return false
	}
	return (digitCnt%2) == 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
