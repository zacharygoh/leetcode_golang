func findErrorNums(nums []int) (result []int) {
    mapping := make(map[int]bool)
    for _, value := range nums {
        if _, found := mapping[value]; found {
            result = append(result, value)
        }else{
            mapping[value] = true
        }
    }
    n := len(nums)
    for n > 0 {
        if !mapping[n] {
            result = append(result, n)
        }
        n--
    }
    return
}

/* 
	Example 1:
	Input: nums = [1,2,2,4]
	Output: [2,3]
*/