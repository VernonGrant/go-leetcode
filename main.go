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

// 3: Given a string s, find the length of the longest substring without
// repeating characters.
//
// The idea was to expand the window to the right until a character is found
// that's withing the current window. At that point we move the left side of the
// window until there's no duplicates matching this character. If there's no
// duplicate we check if the max value needs to be updated.
func lengthOfLongestSubstring(s string) int {
	mx := 0
	r := []rune(s)
	l := len(r)

	if l == 1 {
		return 1
	}

	contains := func(r []rune, c rune) bool {
		for _, sym := range r {
			if sym == c {
				return true
			}
		}
		return false
	}

	for ws, we := 0, 1; we < l; {
		if contains(r[ws:we], r[we]) {
			ws++
		} else {
			if we-ws+1 > mx {
				mx = we - ws + 1
			}
			we++
		}
	}

	return mx
}

func main() {
	// Longest Substring
	fmt.Printf("The lpongest substring is: %d\n", lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Printf("The longest substring is: %d\n", lengthOfLongestSubstring("bbbbb"))     // 1
	fmt.Printf("The longest substring is: %d\n", lengthOfLongestSubstring("pwwkew"))    // 3

	// Max Vowels
	fmt.Printf("Total Vowels: %d\n", maxVowels("weallloveyou", 7)) // 4
	fmt.Printf("Total Vowels: %d\n", maxVowels("abciiidef", 3))    // 3
	fmt.Printf("Total Vowels: %d\n", maxVowels("aeiou", 2))        // 2
	fmt.Printf("Total Vowels: %d\n", maxVowels("leetcode", 3))     // 2
}
