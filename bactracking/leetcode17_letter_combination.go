func letterCombinations(digits string) []string {
    res := []string{}
    if len(digits) == 0 { return res }
    cache := [][]rune {
        {'a','b','c'},
        {'d','e','f'},
        {'g','h','i'},
        {'j','k','l'},
        {'m','n','o'},
        {'p','q','r','s'},
        {'t','u','v'},
        {'w','x','y','z'},
    }
    currRes := []rune{}
    backtrack(0, currRes, digits, cache, &res)
    return res
}

func backtrack(currIdx int, currRes []rune, digits string, cache [][]rune, res *[]string) {
    if len(currRes) == len(digits) { 
        *res = append(*res, string(currRes))
        return
    }
    for _, ch := range cache[digits[currIdx]-48-2] {
        currRes = append(currRes, ch)
        backtrack(currIdx+1, currRes, digits, cache, res)
        currRes = currRes[:len(currRes)-1]
    }
}