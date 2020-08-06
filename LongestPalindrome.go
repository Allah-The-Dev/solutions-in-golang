package main

// func main() {
//   fmt.Println(longestPalindromeString("1234"))
//   fmt.Println(longestPalindromeString("12321"))
//   fmt.Println(longestPalindromeString("9912321456"))
//   fmt.Println(longestPalindromeString("9912333321456"))
//   fmt.Println(longestPalindromeString("12145445499"))
//   fmt.Println(longestPalindromeString("1223213"))
//   fmt.Println(longestPalindromeString("abb"))
// }

func longestPalindromeString(s string) string {
	longest := s[:1]

	for index := range s {
		//for odd length string
		palindrome := getPalindromeForMidRange(s, index, index)
		if len(palindrome) > len(longest) {
			longest = palindrome
		}
		//for even lenght string
		palindrome = getPalindromeForMidRange(s, index, index+1)
		if len(palindrome) > len(longest) {
			longest = palindrome
		}
	}

	return longest
}

func getPalindromeForMidRange(s string, left, right int) string {
	for left >= 0 && right <= (len(s)-1) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
