func sortArrayByParityII(A []int) []int {
    result := make([]int, len(A))
    even, odd := 0, 1
    for _, value := range A {
        if value % 2 == 0 {
            result[even] = value
            even+=2
        }else{
            result[odd] = value
            odd+=2
        }
    }
    return result
}

/* 
	Example 1:

	Input: [4,2,5,7]
	Output: [4,5,2,7]
	Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
*/