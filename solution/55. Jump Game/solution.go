func canJump(nums []int) bool {
    last_index := len(nums)-1
    for i := len(nums)-1; i > -1; i-- {
        if i + nums[i] >= last_index {
            last_index = i
        }
    }
    return last_index == 0
}

/* 
	Greedy solution

	Example 1:

	Input: nums = [2,3,1,1,4]
	Output: true
	Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
	
	Example 2:

	Input: nums = [3,2,1,0,4]
	Output: false
	Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/