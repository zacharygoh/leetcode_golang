func topKFrequent(nums []int, k int) (result []int) {
    if len(nums) == k {
        return nums
    }
    mapping := make(map[int]int)
    for _, value := range nums {
        mapping[value]++
    }
    for ; k > 0; k-- {
        ress, index := 0, -1
        for i, e := range mapping {
            if index < e {
                index, ress = e, i
            }
        }
        delete(mapping, ress)
        result = append(result, ress)
    }
    return
}

/* 
	Example 1:

	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

	Example 2:

	Input: nums = [1], k = 1
	Output: [1]
*/