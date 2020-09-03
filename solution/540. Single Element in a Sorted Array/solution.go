func singleNonDuplicate(nums []int) (result int) {
    if len(nums) == 1 {
        return nums[0]
    }
    if nums[0] != nums[1] {
        return nums[0]
    }
    last := len(nums)-1
    if nums[last] != nums[last-1] {
        return nums[last]
    }
    for i := 1; i < len(nums); i++ {
        if nums[i-1] != nums[i] && nums[i] != nums[i+1] {
            return nums[i]
        }
    }
    return
}

/* 
	Example 1:

	Input: nums = [1,1,2,3,3,4,4,8,8]
	Output: 2

	Example 2:

	Input: nums = [3,3,7,7,10,11,11]
	Output: 10
*/