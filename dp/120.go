package main

func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }
    n := len(triangle)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, i+1)
    }
    dp[0][0] = triangle[0][0]
    for i := 1; i < n; i++ {
        for j := 0; j <= i; j++ {
            if j == 0 {
                dp[i][j] = dp[i-1][j] + triangle[i][j]
            }else if j == i {
                dp[i][j] = dp[i-1][j-1] + triangle[i][j]
            }else {
                if dp[i-1][j-1] < dp[i-1][j] {
                    dp[i][j] = dp[i-1][j-1] + triangle[i][j]
                } else {
                    dp[i][j] = dp[i-1][j] + triangle[i][j]
                }
            }
        }
    }
    minSum := dp[n-1][0]
    for j := 1; j < n; j++ {
        if dp[n-1][j] < minSum {
            minSum = dp[n-1][j]
        }
    }
    
    return minSum
}