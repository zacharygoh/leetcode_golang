/* 
	Kadane's algorithm
*/
func maxSubArray(nums []int) (result int) {
    max_ending := nums[0]
    if len(nums) == 1 {
        return nums[0]
    }
    if nums[0] > 0 {
        result = 0
        max_ending = 0
    }else{
        result = nums[0]
    }
    for _, value := range nums {
        max_ending = max(value, max_ending+value)
        result = max(result, max_ending)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:

	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

	Example 2:

	Input: nums = [1]
	Output: 1

	Example 3:

	Input: nums = [0]
	Output: 0

	Example 4:

	Input: nums = [-1]
	Output: -1

	Example 5:

	Input: nums = [-2147483647]
	Output: -2147483647
*/