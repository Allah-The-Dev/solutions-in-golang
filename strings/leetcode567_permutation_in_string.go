func checkInclusion(s1 string, s2 string) bool {
    // if s1 length is more than s2 return false
    if len(s1) > len(s2) { return false }
    s1Chars := make([]int, 26) // a : 97
    for _, ch := range s1 {
        s1Chars[ch-97] += 1
    }
    // create window of size s1
    s2Chars := make([]int, 26)
    l, r := 0, 0
    for r < len(s2) {
        ch := s2[r]
        s2Chars[ch-97] += 1
        if (r-l+1) == len(s1) {
            if ok := isS1CharsEqualsS2Chars(s1Chars, s2Chars); ok {
                return true
            }
            chL := s2[l]
            s2Chars[chL-97] -= 1
            l++
        }
        r++
    }
    return false
}

func isS1CharsEqualsS2Chars(s1Chars, s2Chars []int) bool {
    for i := 0; i <= 25; i++ {
        if s1Chars[i] != s2Chars[i] { return false }
    }
    return true
}