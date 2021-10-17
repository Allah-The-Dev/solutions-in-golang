// method : bottom up DP
/*
                        12
		   -(1^2)/ -(2^2)|   -(3^2)\
			   11        8           3
since 12 can be solved with min of 1+min(11, 8, 3)

hence it should be solved from bottom to top
i.e. 1->12
*/
func numSquares(n int) int {
    dp := make([]int, n+1)
    // this will set the worst values in dp array
    for i := range dp {
        dp[i] = i
    }
    // start dp from 1->n
    // for each n max dp iteration will be upto sqrt(n)
    for target := range dp {
        if target == 0 { continue }
        sqrtI := int(math.Sqrt(float64(target)))
        for i := 1; i <= sqrtI; i++ {
            square := i*i
            if target - square < 0 {
                break
            }
            dp[target] = min(dp[target], 1+dp[target-square])
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

// top down dfs + memoization
/*
           7
   (7-1^2)/ \(7-2^2)
		 6   3
		/     \
...
*/
func numSquares(n int) int {
    dp := make([]int, n+1)
    // this will set the worst values in dp array
    for i := range dp {
        dp[i] = i
    }
    return dfs(dp, n)
}

func dfs(dp []int, n int) int {
    if n <= 3 { return n }
    if dp[n] != n { return dp[n] }
    sqrtN := int(math.Sqrt(float64(n)))
    for i := 1; i <= sqrtN; i++ {
        dp[n] = min(dp[n], 1+dfs(dp, n-(i*i)))
    } 
    return dp[n]
}

func min(a, b int) int {
    if a < b { return a }
    return b
}