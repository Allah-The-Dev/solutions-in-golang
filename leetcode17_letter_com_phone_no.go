// iterative approach

func letterCombinations(digits string) []string {
    // base case
    if len(digits) == 0 { return []string{} }
    
    cache := [][]rune {
        {'a', 'b', 'c'},
        {'d', 'e', 'f'},
        {'g', 'h', 'i'},
        {'j', 'k', 'l'},
        {'m', 'n', 'o'},
        {'p', 'q', 'r', 's'},
        {'t', 'u', 'v'},
        {'w', 'x', 'y', 'z'},
    }
    
    // 48 - 57 ascii value for 0 - 9
    
    res := []string{}
    
    for i, ch := range digits {
        possibleCharForCurrCh := cache[ch-48-2] // here 2's no value is at 0th index
        if i == 0 {
            for _, currCh := range possibleCharForCurrCh {
                res = append(res, string(currCh))   
            }
            continue
        }
        resLen := len(res)
        for i := 0; i < resLen; i++ {
            currStr := res[0]
            res = res[1:]
            for _, currCh := range possibleCharForCurrCh {
                res = append(res, currStr+string(currCh))   
            }
        }
    }
    return res
}

// recursive approach
func letterCombinations(digits string) []string {
    // base case
    if len(digits) == 0 { return []string{} }
    
    cache := [][]rune {
        {'a', 'b', 'c'},
        {'d', 'e', 'f'},
        {'g', 'h', 'i'},
        {'j', 'k', 'l'},
        {'m', 'n', 'o'},
        {'p', 'q', 'r', 's'},
        {'t', 'u', 'v'},
        {'w', 'x', 'y', 'z'},
    }
    
    // 48 - 57 ascii value for 0 - 9
    
    res := []string{}
    
    backtrack(digits, cache, &res, "", 0)
    
    return res
}

func backtrack(digits string, cache [][]rune, res *[]string, currRes string, digitsIdx int) {
    if digitsIdx == len(digits) { 
        *res  = append(*res, currRes)
        return
    }
    
    currChFromDigit := digits[digitsIdx]
    charsForCurrNo := cache[currChFromDigit-48-2]
    for _, currCh := range charsForCurrNo {
        currRes = currRes+string(currCh)
        backtrack(digits, cache, res, currRes, digitsIdx+1)
        currRes = currRes[:len(currRes)-1]
    }
}