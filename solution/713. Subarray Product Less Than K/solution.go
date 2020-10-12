func numSubarrayProductLessThanK(nums []int, k int) (result int) {
    if k <= 1 {
        return 
    }
    
    left, right, prod := 0, 0, 1
    for right < len(nums) {
        prod *= nums[right]
        for prod >= k {
            prod /= nums[left]
            left++
        }
        result += right - left + 1
        right++
    }
    return
}

/* 
	Example 1:
	
	Input: nums = [10, 5, 2, 6], k = 100
	Output: 8
	Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
	Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
*/