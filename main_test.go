package main

import (
	"reflect"
	"testing"
)

const msg string = "Returned %v, but %v was expected."

func TestMaxVowels(t *testing.T) {
	actual, expected := maxVowels("weallloveyou", 7), 4
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = maxVowels("abciiidef", 3), 3
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = maxVowels("aeiou", 2), 2
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = maxVowels("leetcode", 3), 2
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	actual, expected := lengthOfLongestSubstring("abcabcbb"), 3
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = lengthOfLongestSubstring("bbbbb"), 1
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = lengthOfLongestSubstring("pwwkew"), 3
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}
}

func TestContainsDuplicate(t *testing.T) {
	actual, expected := containsDuplicate([]int{1, 2, 3, 1}), true
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = containsDuplicate([]int{1, 2, 3, 4}), false
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), true
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}
}

func TestIsAnagram(t *testing.T) {
	actual, expected := isAnagram("aacc", "ccac"), false
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = isAnagram("anagram", "nagaram"), true
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = isAnagram("rat", "car"), false
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = isAnagram("anagram", "nagarammm"), false
	if actual != expected {
		t.Errorf(msg, actual, expected)
	}
}

type TestingSlice interface {
	int | string
}

func isEqualSlices[T TestingSlice](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	bubbleSort := func(col []T) {
		l := len(col)
		for i := 0; i < l; i++ {
			for j := 0; j < l - i - 1; j++ {
				if col[j] > col[j+1] {
					col[j], col[j+1] = col[j+1], col[j]
				}
			}
		}
	}

	bubbleSort(a)
	bubbleSort(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestTwoSum(t *testing.T) {
	actual, expected := twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = twoSum([]int{3, 2, 4}, 6), []int{1, 2}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = twoSum([]int{3, 3}, 6), []int{0, 1}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}
}

func TestGroupAnagrams(t *testing.T) {
	actual, expected := groupAnagrams([]string{"", "b"}), [][]string{{""}, {"b"}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = groupAnagrams([]string{"ac", "c"}), [][]string{{"ac"}, {"c"}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = groupAnagrams([]string{""}), [][]string{{""}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = groupAnagrams([]string{"a"}), [][]string{{"a"}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(msg, actual, expected)
	}
}

func TestTopKFrequent(t *testing.T) {
	actual, expected := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), []int{1,2}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = topKFrequent([]int{1}, 1), []int{1}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = topKFrequent([]int{5, 5, 5, 1, 2, 1, 2, 2}, 3), []int{5, 2, 1}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}

	actual, expected = topKFrequent([]int{5, 5, 5, 5, 5}, 1), []int{5}
	if !isEqualSlices(actual, expected) {
		t.Errorf(msg, actual, expected)
	}
}
