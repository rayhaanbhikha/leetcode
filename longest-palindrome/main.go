package main

import "fmt"

type testCase struct {
	input, expect string
}

func main() {

	testCases := []testCase{
		{"babad", "bab"},
		{"bababd", "babab"},
		{"cbbd", "bb"},
		{"bxdabb", "bb"},
	}

	for _, t := range testCases {
		output := longestPalindrome(t.input)
		fmt.Printf("%s -> e: %s a: %s\n", t.input, t.expect, output)
		if output != t.expect {
			fmt.Println("FAIL!", t.input)
			break
		}
	}
}

func longestPalindrome(s string) string {
	n := len(s)
	switch {
	case n == 0, n == 1:
		return s
	case n == 2 && s[0] != s[1]:
		return string(s[0])
	case n == 2:
		return s
	}
	maxPalindromeLen, maxPalindrome, startIndex := 0, "", 0

	for startIndex != n {
		startChar := s[startIndex]
		for i := n - 1; i != startIndex; i-- {
			// 1. distance is greater than maxPalindromeLen then it's worth checking.
			// 2. check if reached repeating char
			if distance := i - startIndex; distance > maxPalindromeLen && s[i] == startChar {
				// 3. check if palindrome
				if potentialPalindrome := s[startIndex : i+1]; isPalindrome(potentialPalindrome) {
					maxPalindromeLen = distance
					maxPalindrome = potentialPalindrome
				}
			}
		}
		startIndex++
	}

	if maxPalindromeLen == 0 {
		return string(s[0])
	}

	return maxPalindrome
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for {
		if l == r || l > r {
			break
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
