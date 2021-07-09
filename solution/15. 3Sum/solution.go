// Two pointers solution
func threeSum(nums []int) (results [][]int) {
    sort.Ints(nums)
    if len(nums) == 0 || len(nums) < 3 {
        return
    }
    n := len(nums) - 1
    for i := 0; i < n; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        left, right := i + 1, n
        for left < right {
            if nums[left] + nums[right] > -nums[i] {
                right--
            } else if nums[left] + nums[right] < -nums[i] {
                left++
            } else {
                results = append(results, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left + 1] {
                    left++
                }
                for left < right && nums[right] == nums[right - 1] {
                    right--
                }
                left++
                right--
            }
        }
    }
    return
}

/* 
	Example 1:
	Input: nums = [-1,0,1,2,-1,-4]
	Output: [[-1,-1,2],[-1,0,1]]

	Example 2:
	Input: nums = []
	Output: []

	Example 3:
	Input: nums = [0]
	Output: []
*/