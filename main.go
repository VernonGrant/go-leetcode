package main

import (
	"fmt"
)

///////////
// Other //
///////////

// 1456: Maximum number of vowels in a substring of given length.  Given a
// string s and an integer k, return the maximum number of vowel letters in any
// substring of s with length k.  Vowel letters in English are 'a', 'e', 'i',
// 'o', and 'u'.
//
// (Sliding Window)
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
//
// (Sliding Window)
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

////////////////////////
// Arrays and Hashing //
////////////////////////

// 217: Given an integer array nums, return true if any value appears at least
// twice in the array, and return false if every element is distinct.
func containsDuplicate(nums []int) bool {
	checked := make(map[int]bool)
	l := len(nums)

	for i := 0; i < l; i++ {
		_, ok := checked[nums[i]]
		if ok {
			return true
		} else {
			checked[nums[i]] = true
		}
	}

	return false
}

// 242: Given two strings s and t, return true if t is an anagram of s, and
// false otherwise.  An Anagram is a word or phrase formed by rearranging the
// letters of a different word or phrase, typically using all the original
// letters exactly once.
//
// TIP: We can also sort the second string and compare the sorted results.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	strRunes := make(map[rune]int)
	for _, v := range s {
		strRunes[v]++
	}
	for _, v := range t {
		strRunes[v]--
	}
	for k, _ := range strRunes {
		if strRunes[k] != 0 {
			return false
		}
	}

	return true
}

// 1: Given an array of integers nums and an integer target, return indices of
// the two numbers such that they add up to target. You may assume that each
// input would have exactly one solution, and you may not use the same element
// twice.  You can return the answer in any order.
//
// This is making us of complements, we basically calculate target - v, and
// check if the result exists within a hash map that contains the values and
// indexes of array elements.
func twoSum(nums []int, target int) []int {
	comp := make(map[int]int)

	for i, v := range nums {
		idx, ok := comp[target-v]
		if ok {
			return []int{idx, i}
		} else {
			comp[v] = i
		}
	}

	return []int{0, 0}
}

func main() {
	// TODO: Add assertion helper function.

	// Two Sum
	fmt.Printf("The two sum is %v\n", twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Printf("The two sum is %v\n", twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Printf("The two sum is %v\n", twoSum([]int{3, 3}, 6))         // [0,1]
	fmt.Println("---")

	// Valid Anagram
	fmt.Printf("Is Anagram: %t\n", isAnagram("aacc", "ccac"))         // false
	fmt.Printf("Is Anagram: %t\n", isAnagram("anagram", "nagaram"))   // true
	fmt.Printf("Is Anagram: %t\n", isAnagram("rat", "car"))           // false
	fmt.Printf("Is Anagram: %t\n", isAnagram("anagram", "nagarammm")) // true
	fmt.Println("---")

	// Contains duplicate
	fmt.Printf("Contains duplicate: %t\n", containsDuplicate([]int{1, 2, 3, 1}))                   // true
	fmt.Printf("Contains duplicate: %t\n", containsDuplicate([]int{1, 2, 3, 4}))                   // false
	fmt.Printf("Contains duplicate: %t\n", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true
	fmt.Println("---")

	// Longest Substring
	fmt.Printf("The longest substring is: %d\n", lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Printf("The longest substring is: %d\n", lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Printf("The longest substring is: %d\n", lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println("---")

	// Max Vowels
	fmt.Printf("Total Vowels: %d\n", maxVowels("weallloveyou", 7)) // 4
	fmt.Printf("Total Vowels: %d\n", maxVowels("abciiidef", 3))    // 3
	fmt.Printf("Total Vowels: %d\n", maxVowels("aeiou", 2))        // 2
	fmt.Printf("Total Vowels: %d\n", maxVowels("leetcode", 3))     // 2
	fmt.Println("---")
}
