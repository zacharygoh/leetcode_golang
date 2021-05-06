func minCostClimbingStairs(cost []int) (result int) {
    last := len(cost)
    for i := 2; i < last; i++ {
        cost[i] += min(cost[i-1], cost[i-2])
    }
	return min(cost[last-1], cost[last-2])
	/* 
		Extra space
	*/
    // last := len(cost)
    // dp := make([]int, last)
    // dp[0], dp[1] = cost[0], cost[1]
    // for i := 2; i < last; i++ {
    //     dp[i] = cost[i] + min(dp[i-1], dp[i-2])
    // }
    // return min(dp[last-1], dp[last-2])
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}