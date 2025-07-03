package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Preprocess the string to handle even-length palindromes
	t := preprocess(s)
	n := len(t)
	p := make([]int, n)
	center, right := 0, 0
	maxLen := 0
	centerIndex := 0

	for i := 1; i < n-1; i++ {
		mirror := 2*center - i

		if i < right {
			p[i] = min(right-i, p[mirror])
		}

		// Expand around center i
		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}

		// Update center and right boundary
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}

		// Track the longest palindrome
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	// Get the start index in the original string
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

// Insert # between characters to handle even/odd symmetry
func preprocess(s string) string {
	if len(s) == 0 {
		return "^$"
	}
	res := "^"
	for _, ch := range s {
		res += "#" + string(ch)
	}
	res += "#$"
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindrome("babad"))   // "bab" or "aba"
	fmt.Println(longestPalindrome("cbbd"))    // "bb"
	fmt.Println(longestPalindrome("racecar")) // "racecar"
}
