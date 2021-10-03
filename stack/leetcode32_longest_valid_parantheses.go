func longestValidParentheses(s string) int {
    res := 0
    stack := []int{-1}
    for i, ch := range s {
        if ch == '(' {
            stack = append(stack, i)
        } else {
            stack  = stack[:len(stack)-1]
            if len(stack) == 0 { 
                stack = append(stack, i)
            }
            stackTop := stack[len(stack)-1]
            res = max(res, i-stackTop)
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}