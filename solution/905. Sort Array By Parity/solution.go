func sortArrayByParity(A []int) []int {
    result := make([]int, len(A))
    start, end := 0, len(A)-1
    for _, value := range A {
        if value % 2 == 0 {
            result[start] = value
            start++
        }else{
            result[end] = value
            end--
        }
    }
    return result
}

/* 
	Example 1:

	Input: [3,1,2,4]
	Output: [2,4,3,1]
	The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
*/