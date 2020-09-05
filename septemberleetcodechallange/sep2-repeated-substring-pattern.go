package septemberleetcodechallange

import (
	"strings"
)

func RepeatedSubstringPattern(s string) bool {
	var curStr strings.Builder
	var repeatStr strings.Builder
	for i, ch := range s {
		curStr.WriteRune(ch)
		repeatStr.Reset()
		repeatStr.WriteString(curStr.String())
		for {
			repeatStr.WriteString(curStr.String())
			if repeatStr.Len() >= len(s) {
				break
			}
		}
		if repeatStr.String() == s {
			return true
		}
		if i == ((len(s) / 2) - 1) {
			break
		}
	}
	return false
}
