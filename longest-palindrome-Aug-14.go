package main

import (
	"fmt"
	"unicode"
)

func longestPalindrome(s string) int {
	charCountArr := make([]int, 52)
	for _, ch := range s {
		if unicode.IsLower(ch) {
			charCountArr[int(ch)-97]++
		} else {
			charCountArr[int(ch)-38]++
		}
	}
	result := 0
	oddIsPresent := false
	for _, count := range charCountArr {
		if count%2 == 0 {
			result += count
		} else {
			result += count - 1
			oddIsPresent = true
		}
	}
	fmt.Printf("%v\n", charCountArr)
	if oddIsPresent {
		return result + 1
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
