func deleteAndEarn(nums []int) (result int) {
    if len(nums) == 0 {
        return
    }
    r := 0
    for _, value := range nums {
        r = max(r, value)
    }
    points := make([]int, r + 1)
    for _, value := range nums {
        points[value] += value
    }
    return rob(points)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) (result int) {
    dp2 := 0
    for _, value := range nums {
        dp := max(dp2 + value, result)
        dp2 = result
        result = dp
    }
    return
}

/* 
	Example 1:

	Input: nums = [3, 4, 2]
	Output: 6
	Explanation: 
	Delete 4 to earn 4 points, consequently 3 is also deleted.
	Then, delete 2 to earn 2 points. 6 total points are earned.
	

	Example 2:

	Input: nums = [2, 2, 3, 3, 3, 4]
	Output: 9
	Explanation: 
	Delete 3 to earn 3 points, deleting both 2's and the 4.
	Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
	9 total points are earned.`
*/