func coinChange(coins []int, amount int) (coins_required int) {
    if len(coins) == 0 {
        return -1
    }
    if amount == 0 {
        return 0
    }
    dp := make([]int, amount + 1)
    for i := 1; i <= amount; i++ {
        min_coin := math.MaxInt32
        for _, coin := range coins {
            if i - coin >= 0 && dp[i-coin] != -1 {
                min_coin = min(min_coin, dp[i-coin]+1)
            }
        }
        
        if min_coin == math.MaxInt32 {
            dp[i] = -1
        }else{
            dp[i] = min_coin
        }
    }
    return dp[amount]
    /* Greedy solution doens't work */
    // sort.Sort(sort.Reverse(sort.IntSlice(coins)))
    // for _, coin := range coins {
    //     for n >= coin {
    //         // n = n - coin
    //         // coins_required += 1
    //     }
    // }
    // if n != 0 {
    //     return -1
    // }
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

/* 
Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3 
Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3
Output: -1
*/