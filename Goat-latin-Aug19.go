package main

import (
	"strings"
	"unicode"
)

func toGoatLatin(S string) string {
	stringsArr := strings.Fields(S)
	var aStr strings.Builder
	aStr.WriteString("a")
	for i, str := range stringsArr {
		var b strings.Builder
		switch strRuneArr := []rune(str); unicode.ToLower(strRuneArr[0]) {
		case 'a', 'e', 'i', 'o', 'u':
			b.WriteString(str)
			b.WriteString("ma")
			b.WriteString(aStr.String())
			stringsArr[i] = b.String()
			b.Reset()
			aStr.WriteString("a")
		default:
			b.WriteString(string(strRuneArr[1:]))
			b.WriteRune(strRuneArr[0])
			b.WriteString("ma")
			b.WriteString(aStr.String())
			stringsArr[i] = b.String()
			b.Reset()
			aStr.WriteString("a")
		}
	}
	return strings.Join(stringsArr, " ")
}

// func main() {
// 	fmt.Println(toGoatLatin("I speak Goat Latin"))
// }
