package main

import "fmt"

func main() {
	fmt.Println("abcabcbb", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("nfpdmpi", lengthOfLongestSubstring("nfpdmpi"))
	fmt.Println(" ", lengthOfLongestSubstring(" "))
	fmt.Println("aab", lengthOfLongestSubstring("aab"))
	fmt.Println("bbbbbbb", lengthOfLongestSubstring("bbbbbbb"))
	fmt.Println("dvdf", lengthOfLongestSubstring("dvdf"))
}

// sliding window implementation
func lengthOfLongestSubstring(s string) int {
	if x := len(s); x == 1 || x == 0 {
		return x
	}

	maxSubStringlen := 0
	n := len(s)
	startIndex := 0

	for startIndex != n {
		stored := make(map[rune]struct{})
		subStringlen := 0
		for i := startIndex; i < n; i++ {
			c := rune(s[i])
			if _, ok := stored[c]; !ok {
				stored[c] = struct{}{}
				subStringlen++
			} else {
				startIndex++
				break
			}
			if i == n-1 {
				startIndex++
			}
			maxSubStringlen = max(subStringlen, maxSubStringlen)
		}

	}

	return maxSubStringlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// simple while loop implmentation

// func lengthOfLongestSubstring(s string) int {
// 	maxSubStringlen := 0

// 	for len(s) != 0 {
// 		stored := make(map[rune]struct{})
// 		n := len(s)
// 		subStringlen := 0
// 		for index, c := range s {
// 			if _, ok := stored[c]; !ok {
// 				stored[c] = struct{}{}
// 				subStringlen++
// 			} else {
// 				fmt.Println(s, subStringlen, maxSubStringlen, index, n-1)

// 				if subStringlen > maxSubStringlen {
// 					maxSubStringlen = subStringlen
// 				}

// 				if index == n-1 {
// 					return maxSubStringlen
// 				}
// 				break
// 			}
// 		}
// 		if subStringlen > maxSubStringlen {
// 			maxSubStringlen = subStringlen
// 		}
// 		s = s[1:]
// 	}

// 	return maxSubStringlen
// }

// recusive implementation

// func lengthOfLongestSubstring(s string) int {
// 	return l(s, 0)
// }

// func l(s string, maxSubStringlen int) int {
// 	stored := make(map[rune]struct{})
// 	n := len(s)
// 	subStringlen := 0
// 	for index, c := range s {
// 		if _, ok := stored[c]; !ok {
// 			stored[c] = struct{}{}
// 			subStringlen++
// 		} else {
// 			fmt.Println(s, subStringlen, maxSubStringlen, index, n-1)

// 			if subStringlen < maxSubStringlen {
// 				subStringlen = maxSubStringlen
// 			}

// 			if index == n-1 {
// 				return subStringlen
// 			}
// 			return l(s[1:], subStringlen)
// 		}
// 	}
// 	if subStringlen < maxSubStringlen {
// 		subStringlen = maxSubStringlen
// 	}
// 	return subStringlen
// }
