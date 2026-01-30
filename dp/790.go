package main

func numTilings(n int) int {
    const mod = 1000000007

    if n <= 2 {
        return n
    }
    if n == 3 {
        return 5
    }

    dp := make([]int64, n+1)
    dp[0] = 1
    dp[1] = 1
    dp[2] = 2
    dp[3] = 5

    for i := 4; i <= n; i++ {
        dp[i] = (2*dp[i-1] + dp[i-3]) % mod
    }

    return int(dp[n])
}
