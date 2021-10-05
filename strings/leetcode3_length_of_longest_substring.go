// 1. use cache to store last occurence index of a char 
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 { return 0 }
    cache := make(map[byte]int)
    l, r := 0, 0
    res := 0
    for r < len(s) {
        if idx, ok := cache[s[r]]; ok {
            l = max(l, idx+1)
        }
        cache[s[r]] = r
        res = max(res, r-l+1)
        r++
    }
    return res
}

func max(a, b int) int {
    if a > b { return a }
    return b
}