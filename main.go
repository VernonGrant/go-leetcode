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

// 9: Given an integer x, return true if x is a palindrome, and false otherwise.
//
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it
// becomes 121-. Therefore it is not a palindrome.
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
// Constraints:
// -231 <= x <= 231 - 1
func isPalindrome(x int) bool {

	return false
}

// 14: Write a function to find the longest common prefix string amongst an
// array of strings.  If there is no common prefix, return an empty string "".
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// Constraints:
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.
func longestCommonPrefix(strs []string) string {

	return ""
}

// 20: Given a string s containing just the characters '(', ')', '{', '}', '[' and
// ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
// Input: s = "()"
// Output: true
//
// Example 2:
// Input: s = "()[]{}"
// Output: true
//
// Example 3:
// Input: s = "(]"
// Output: false
func isValidParentheses(s string) bool {

	return false
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

// 49: Given an array of strings strs, group the anagrams together. You can
// return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
// different word or phrase, typically using all the original letters exactly
// once.
func groupAnagrams(strs []string) [][]string {
	wordFreq := func(w string) [27]int {
		var alpha [27]int
		l := len(w)
		for i := 0; i < l; i++ {
			alpha[int(w[i])-int('a')]++
		}
		return alpha
	}

	master := [][]string{}
	matched := make(map[string]bool)
	l := len(strs)

	for i := 0; i < l; i++ {
		group := []string{strs[i]}
		groupFreq := wordFreq(group[0])

		_, ok := matched[strs[i]]
		if ok {
			continue
		}

		for j := i + 1; j < l; j++ {
			if groupFreq == wordFreq(strs[j]) {
				matched[strs[j]] = true
				group = append(group, strs[j])
			}
		}

		master = append(master, group)
	}

	return master
}

// 347: Given an integer array nums and an integer k, return the k most frequent
// elements. You may return the answer in any order.
//
// NOTE: Using an inverse bucket sort can solve this in a an O(n) time. By
// storing the counts as indexes (based on k) and the values inside
// buckets. Using an array.
//
// TODO: Implement bucket sort version.
func topKFrequent(nums []int, k int) []int {
	// Calculate the frequencies of each number.
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Get the highest frequencies.
	result := []int{}
	checked := make(map[int]bool)
	for i := 0; i < k; i++ {
		max := [2]int{}

		for k, v := range freq {
			if v >= max[1] {
				_, ok := checked[k]
				if !ok {
					max[0] = k
					max[1] = v
				}
			}
		}

		checked[max[0]] = true
		result = append(result, max[0])
	}

	return result
}

// 238: Given an integer array nums, return an array answer such that answer[i] is
// equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
// integer.
//
// You must write an algorithm that runs in O(n) time and without using the
// division operation.
//
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
//
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
func productExceptSelf(nums []int) []int {
	// Posfix and Prefix solution.

	// Calculate prefixes
	// Calculate suffixes

	// 1 + 2 + 3 + 4 = 10
	// 1
	// No division can be used.
	// 2 * 3 * 4 = 24
	// Can we get the two resulting arrays and sum them.

	// O(n)

	// go through the loop only once?

	return []int{}
}

// 36: Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to
// be validated according to the following rules:
//
// Each row must contain the digits 1-9 without repetition.
//
// Each column must contain the digits 1-9 without repetition.
//
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
// without repetition.
//
// A Sudoku board (partially filled) could be valid but is not necessarily
// solvable.  Only the filled cells need to be validated according to the
// mentioned rules.
//
// Input: board =
//
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
//
// Output: true
//
// Input: board =
//
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
//
// Output: false Explanation: Same as Example 1, except with the 5 in the top
// left corner being modified to 8. Since there are two 8's in the top left 3x3
// sub-box, it is invalid.
func isValidSudoku(board [][]byte) bool {

	return false
}

// Given an unsorted array of integers nums, return the length of the longest
// consecutive elements sequence.  You must write an algorithm that runs in O(n)
// time.
//
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3,
// 4]. Therefore its length is 4.
//
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
func longestConsecutive(nums []int) int {

	return 0
}

//////////////////
// Linked Lists //
//////////////////

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

// 2: You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit.
//
// Add the two numbers and return the sum as a linked list.  You may assume the
// two numbers do not contain any leading zero, except the number 0 itself.
//

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Input: l1 = [0], l2 = [0]
// Output: [0]

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

func addTwoNumbersLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var temp ListNode;
	return &temp
}

// 21: You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists in a one sorted list. The list should be made by splicing
// together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:
//
// Input: list1 = [], list2 = []
// Output: []
// Example 3:
//
// Input: list1 = [], list2 = [0]
// Output: [0]
//
// Constraints:
//
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var temp ListNode;
	return &temp
}

func main() {
	fmt.Println("Hello from main.")
}
