func decompressRLElist(nums []int) (result []int) {
    if len(nums) % 2 != 0 {
        return 
    }
    
    for i := 1; i < len(nums); i+=2 {
        for j := 0; j < nums[i-1]; j++ {
            result = append(result, nums[i])
        }
    }
    return
}

/* 
	Example 1:

	Input: nums = [1,2,3,4]
	Output: [2,4,4,4]
	Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
	The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
	At the end the concatenation [2] + [4,4,4] is [2,4,4,4].

	Example 2:

	Input: nums = [1,1,2,3]
	Output: [1,3,3]
*/