package main

import (
	"fmt"
)

// 1456: Maximum number of vowels in a substring of given length.  Given a
// string s and an integer k, return the maximum number of vowel letters in any
// substring of s with length k.  Vowel letters in English are 'a', 'e', 'i',
// 'o', and 'u'.
//
// TODO: There's a lot of performance to gain in string rune access. Run some
// benchmarks for string rune access operations.
func maxVowels(s string, k int) int {
	isVowel := func(r rune) bool {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	runes := []rune(s)
	length := len(runes)

	// Calculate initial vowel count.
	cnt := 0
	for i := 0; i < k; i++ {
		if isVowel(runes[i]) {
			cnt++
		}
	}

	// Handle window position changes.
	mx := cnt
	for ws, we := 1, k; we < length; {
		if isVowel(runes[ws-1]) {
			cnt--
		}
		if isVowel(runes[we]) {
			cnt++
		}
		if cnt > mx {
			mx = cnt
		}
		ws, we = ws+1, we+1
	}

	return mx
}

func main() {
	// Max Vowels
	fmt.Printf("Total Vowels: %d\n", maxVowels("weallloveyou", 7)) // 4
	fmt.Printf("Total Vowels: %d\n", maxVowels("abciiidef", 3))    // 3
	fmt.Printf("Total Vowels: %d\n", maxVowels("aeiou", 2))        // 2
	fmt.Printf("Total Vowels: %d\n", maxVowels("leetcode", 3))     // 2
}
