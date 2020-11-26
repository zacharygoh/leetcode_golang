func videoStitching(clips [][]int, T int) (result int) {
    max := 101
    dp := make([][]int, T+1)
    for i, _ := range dp {
        dp[i] = make([]int, T+1)
        for ii, _ := range dp[i] {
            dp[i][ii] = max
        }
    }
    for _, value := range clips {
        start, end := value[0], value[1]
        for i := 1; i <= T; i++ {
            for j := 0; j <= T - i; j++ {
                k := i + j
                if start > k || end < j {
                    //out of bound
                    continue
                }
                if start <= j && end >= k {
                    //start or end
                    dp[j][k] = 1
                }else if end >= k {
                    //partially covered
                    //start in between, end after close
                    dp[j][k] = Min(dp[j][k], dp[j][start] + 1)
                }else if start <= j {
                    //partially covered
                    //start early, end before close
                    dp[j][k] = Min(dp[j][k], dp[end][k] + 1)
                }else{
                    //cover up the uncovered area
                    dp[j][k] = Min(dp[j][k], dp[j][start] + 1 + dp[end][k])
                }
            }
        }
    }
    if dp[0][T] == max {
        return -1
    }else{
        return dp[0][T]
    }
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

/* 
	Example 1:

	Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
	Output: 3
	Explanation: 
	We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
	Then, we can reconstruct the sporting event as follows:
	We cut [1,9] into segments [1,2] + [2,8] + [8,9].
	Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].

	Example 2:

	Input: clips = [[0,1],[1,2]], T = 5
	Output: -1
	Explanation: 
	We can't cover [0,5] with only [0,1] and [1,2].

	Example 3:

	Input: clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]], T = 9
	Output: 3
	Explanation: 
	We can take clips [0,4], [4,7], and [6,9].

	Example 4:

	Input: clips = [[0,4],[2,8]], T = 5
	Output: 2
	Explanation: 
	Notice you can have extra video after the event ends.
*/