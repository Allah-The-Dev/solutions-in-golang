/*
Given a string s, reverse the order of characters in each word 
within a sentence while still preserving whitespace and initial word order.

Example 1:
----------
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

*/
func reverseWords(s string) string {
    sBytes := []byte(s)
    sLen := len(s)
    whitespaceAsciiCode := byte(32)
    l, r := 0, 0
    for i, ch := range sBytes {
        if ch == whitespaceAsciiCode || i == sLen-1 {
            r = i-1 // if this is not the last word
            if i == sLen-1 { r = i }// if this is the last word
            reverse(sBytes, l, r)
            if i+1 < sLen {
                l = i+1
            }
        }
    }
    return string(sBytes)
}

func reverse(sBytes []byte, l, r int) {
    for ;l < r; l, r = l+1, r-1 {
        sBytes[l], sBytes[r] = sBytes[r], sBytes[l]
    }
}