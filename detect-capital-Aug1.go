package main

import (
	"unicode"
)

func detectCapitalUse(word string) bool {
	otherCharIsUpper := false
	firstCharIsLower := false
	for i, c := range word {
		if i == 0 {
			firstCharIsLower = unicode.IsLower(c)
			continue
		}
		if firstCharIsLower && unicode.IsUpper(c) {
			return false
		}
		if i == 1 {
			otherCharIsUpper = unicode.IsUpper(c)
		}
		if unicode.IsUpper(c) && !otherCharIsUpper {
			return false
		}
		if unicode.IsLower(c) && otherCharIsUpper {
			return false
		}
		if unicode.IsUpper(c) {
			otherCharIsUpper = true
		}
	}
	return true
}

// func main() {
// 	fmt.Println(detectCapitalUse("Aaaa"))
// }
