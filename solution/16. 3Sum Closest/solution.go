func solution(nums []int, target int) (result int) {
	sort.Ints(nums)
	result = math.MinInt32
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			tempNum := nums[i] + nums[j] + nums[k]

			// Try to find the closest index
			if tempNum < target {
				j++
			} else if tempNum > target {
				k--
			} else {
				// If matched then return
				return target
			}

			// Compare to get the closest result
			if abs(target-tempNum) < abs(target-result) {
				result = tempNum
			}
		}
	}
	return
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

/*
	Example 1:
	Input: nums = [-1,2,1,-4], target = 1
	Output: 2
	Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

	Example 2:
	Input: nums = [0,0,0], target = 1
	Output: 0
*/