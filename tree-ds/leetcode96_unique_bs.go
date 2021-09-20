func numTrees(n int) int {
    
    if n == 1 { return 1}
    
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    
    return calculateCatlanNo(n, dp)
}

func calculateCatlanNo(n int, dp []int ) int {
    
    for i := 2; i <= n; i++ {
        for j := 0; j < i; j++ {
            dp[i] += dp[j]*dp[i-j-1]
        }
    }
    
    return dp[n]
}

/*func calculateCatlanNo(n int, dp []int ) int {
    if dp[n] != 0 { return dp[n] }
    
    for i := 0; i < n; i++ {
        dp[n] += calculateCatlanNo(i, dp)*calculateCatlanNo(n-i-1, dp)
    }
    
    return dp[n]
}
*/
