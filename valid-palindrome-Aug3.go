package main

import (
	"unicode"
)

func isPalindrome(s string) bool {

	sRune := []rune(s)

	if len(s) == 1 {
		return true
	}

	for i, j := 0, len(s)-1; i <= j; {
		if !isAlphaNumeric(sRune[i]) {
			i++
			continue
		} else if !isAlphaNumeric(sRune[j]) {
			j--
			continue
		}
		if unicode.IsDigit(sRune[i]) && unicode.IsLetter(sRune[j]) {
			return false
		} else if unicode.IsDigit(sRune[j]) && unicode.IsLetter(sRune[i]) {
			return false
		} else if unicode.IsDigit(sRune[i]) && unicode.IsDigit(sRune[j]) && sRune[i] != sRune[j] {
			return false
		} else if unicode.IsLetter(sRune[i]) && unicode.IsLetter(sRune[j]) && unicode.ToLower(sRune[i]) != unicode.ToLower(sRune[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaNumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

// func main() {
// 	fmt.Println(isPalindrome(".,"))
// }
