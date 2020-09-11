func longestCommonSubsequence (text1 string, text2 string) int {
    dp := make([][]int, len(text1))
    for i:=0; i < len(text1); i++ {
        dp[i] = make([]int, len(text2))
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = -1
        }
    }
    return lcs(text1, text2, len(text1), len(text2), dp)
}

func lcs (text1, text2 string, x, y int, dp [][]int) int {
    if x == 0 || y == 0 {
        return 0
    }
    if dp[x-1][y-1] != -1{
        return dp[x-1][y-1]
    }
    if text1[x-1] == text2[y-1] {
        dp[x-1][y-1] = 1 + lcs(text1, text2, x-1, y-1, dp)
        return dp[x-1][y-1]
    }else{
        dp[x-1][y-1] = Max(lcs(text1, text2, x, y-1, dp), lcs(text1, text2, x-1, y, dp))
        return dp[x-1][y-1]
    }
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

/* 
	Example 1:

	Input: text1 = "abcde", text2 = "ace" 
	Output: 3  
	Explanation: The longest common subsequence is "ace" and its length is 3.

	Example 2:

	Input: text1 = "abc", text2 = "abc"
	Output: 3
	Explanation: The longest common subsequence is "abc" and its length is 3.

	Example 3:

	Input: text1 = "abc", text2 = "def"
	Output: 0
	Explanation: There is no such common subsequence, so the result is 0.
*/