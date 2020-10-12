func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        median := (low + high)/2
        if nums[median] == target {
            return median
        }else if nums[median] < target {
            low = median + 1
        }else{
            high = median - 1
        }
    }
    return -1
}

/* 
	Example 1:

	Input: nums = [-1,0,3,5,9,12], target = 9
	Output: 4
	Explanation: 9 exists in nums and its index is 4

	Example 2:

	Input: nums = [-1,0,3,5,9,12], target = 2
	Output: -1
	Explanation: 2 does not exist in nums so return -1
*/