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

// sliding window implementation
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}

	startIndex := 0
	maxSubstringLen := 0
	
	for startIndex != n {
		store := make(map[byte]struct{})
		substring := ""
		substringLen := 0
		breakout := false
		for i := startIndex; i < n; i++ {
			if i == n-1 {
				breakout = true
			}
			c := s[i]
			//1. check if in store
			if _, ok := store[c]; !ok {
				// 2. if not then add
				store[c] = struct{}{}
				// 3. add to substring and increment substringlen
				substring += string(c)
				substringLen++
				if substringLen > maxSubstringLen {
					maxSubstringLen = substringLen
				}
			} else {
				// 2. already in store
				// 2a. record substring len.
				// 2b. update store.
				startIndex++
				if i != n-1 {
					break
				}
			}
		}
		if breakout {
			break
		}
	}

	return maxSubstringLen
}
