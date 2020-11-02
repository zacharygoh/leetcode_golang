func sortedSquares(A []int) (result []int) {
    for _, value := range A {
        result = append(result, value*value)
    }
    sort.Ints(result)
    return
}

/* 
	Example 1:

	Input: [-4,-1,0,3,10]
	Output: [0,1,9,16,100]

	Example 2:

	Input: [-7,-3,2,3,11]
	Output: [4,9,9,49,121]
*/