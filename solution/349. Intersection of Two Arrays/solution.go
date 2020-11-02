func intersection(nums1 []int, nums2 []int) (result []int) {
    mapping := make(map[int]bool)
    for _, value := range nums1 {
        mapping[value] = true
    }
    for _, value := range nums2 {
        if mapping[value] {
            mapping[value] = false
            result = append(result, value)
        }
    }
    return
}


/* 
	Example 1:

	Input: nums1 = [1,2,2,1], nums2 = [2,2]
	Output: [2]

	Example 2:

	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	Output: [9,4]
*/