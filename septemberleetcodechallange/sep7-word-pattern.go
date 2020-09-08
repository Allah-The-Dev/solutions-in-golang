package septemberleetcodechallange

import "strings"

func WordPattern(pattern string, str string) bool {
	strFields := strings.Fields(str)
	patternToStrMap := make(map[rune]string)
	strToPatternMap := make(map[string]rune)
	if len(strFields) != len(pattern) {
		return false
	}
	//creating a map of string to pattern characters in strFields
	for i, ch := range pattern {
		ithStringInStrFields := strFields[i]
		if value, ok := patternToStrMap[ch]; ok && value != ithStringInStrFields {
			return false
		} else {
			patternToStrMap[ch] = ithStringInStrFields
		}
		if value, ok := strToPatternMap[ithStringInStrFields]; ok && value != ch {
			return false
		} else {
			strToPatternMap[ithStringInStrFields] = ch
		}
	}
	return true
}
