func findMin(nums []int) (result int) {
	/* 
		O(n)
		first = nums[0]
		for _, value := range nums {
			if first > value {
				first = value
			}
		}
		return
	*/

    if len(nums) == 1 {
        return nums[0]

    if len(nums) == 2 {
        if nums[0] > nums[1] {
            return nums[1]
        }
        return nums[0]
    }
    
    pivot_point := len(nums)/2
    is_right := true
    if nums[pivot_point+1] > nums[pivot_point-1] {
        is_right = false
    }
    
    result = nums[pivot_point]
    if is_right {
        for i := pivot_point; i < len(nums); i++ {
            if result > nums[i] {
                result = nums[i]
            }
        }
    }else{
        for i := pivot_point; i > -1; i-- {
            if result > nums[i] {
                result = nums[i]
            }
        }
    }
    if nums[0] < result {
        return nums[0]
    }
    if nums[len(nums)-1] < result {
        for i := len(nums)-1; i > pivot_point; i-- {
            if nums[i] < result {
                result = nums[i]
            }else{
                break
            }
        }
    }
	return
}

/* 
	Example 1:

	Input: [3,4,5,1,2] 
	Output: 1

	Example 2:

	Input: [4,5,6,7,0,1,2]
	Output: 0
*/