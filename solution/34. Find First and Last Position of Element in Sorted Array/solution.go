// Bruteforce method
func searchRange(nums []int, target int) (result []int) {
    mapping := make(map[int][]int)
    for index, value := range nums {
        if value == target {
            mapping[target] = append(mapping[target], index)
        }
    }
    for _, value := range mapping {
        if len(value) > 0 {
            result = append(result, value[0])
            result = append(result, value[len(value) - 1])
            break
        }
    }
    if len(result) == 0 {
        result = append(result, -1)
        result = append(result, -1)
    }
    return
}

/* 
	Example 1:

	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]

	Example 2:

	Input: nums = [5,7,7,8,8,10], target = 6
	Output: [-1,-1]

	Example 3:

	Input: nums = [], target = 0
	Output: [-1,-1]
*/