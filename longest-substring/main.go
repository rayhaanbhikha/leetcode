package main

import "fmt"

type testCase struct {
	input  string
	expect int
}

func main() {
	testCases := []testCase{
		{"", 0},
		{"a", 1},
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"nfpdmpi", 5},
		{"aab", 2},
		{"bbbbbbb", 1},
		{"dvdf", 3},
		{"anviaj", 5},
	}

	for _, t := range testCases {
		output := lengthOfLongestSubstring(t.input)
		fmt.Printf("%s -> e: %d a: %d\n", t.input, t.expect, output)
		if output != t.expect {
			fmt.Println("FAIL!", t.input)
			break
		}
	}

}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}

	store := make(map[byte]int)
	maxSubstringLen := 0
	substringLen := 0

	for i := 0; i < n; i++ {
		c := s[i]
		charIdx, ok := store[c]
		if !ok {
			store[c] = i
			substringLen++
			maxSubstringLen = max(substringLen, maxSubstringLen)
		} else {
			store[c] = i
			for k, val := range store {
				if val <= charIdx {
					delete(store, k)
					substringLen--
				}
			}
		}
	}

	return maxSubstringLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
