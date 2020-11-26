func uniquePaths(m int, n int) int {
    if m == 1 && n == 1 {
        return 1
    }
    dp := make([][]int, m)
    for i, _ := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    for i, _ := range dp[0] {
        dp[0][i] = 1
    }
    for i := 1; i < m; i++ {
        for k := 1; k < n; k++ {
            dp[i][k] = dp[i-1][k] + dp[i][k-1]
        }
    }
    return dp[m-1][n-1]
}

/* 
	Example 1:

	Input: m = 3, n = 7
	Output: 28

	Example 2:

	Input: m = 3, n = 2

	Output: 3
	Explanation:
	From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
	1. Right -> Down -> Down
	2. Down -> Down -> Right
	3. Down -> Right -> Down

	Example 3:

	Input: m = 7, n = 3
	Output: 28

	Example 4:

	Input: m = 3, n = 3
	Output: 6
*/