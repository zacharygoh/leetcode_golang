func smallestRangeI(A []int, K int) (result int) {
    min, max := A[0], A[0]
    for _, value := range A {
        if min > value {
            min = value 
        }
        if max < value {
            max = value
        }
    }
    fmt.Println(max, min, K)
    if (max - K) - (min + K) > 0 {
        return (max - K) - (min + K)
    }else{
        return
    }
}

/* 
	Example 1:

	Input: A = [1], K = 0
	Output: 0
	Explanation: B = [1]

	Example 2:

	Input: A = [0,10], K = 2
	Output: 6
	Explanation: B = [2,8]

	Example 3:

	Input: A = [1,3,6], K = 3
	Output: 0
	Explanation: B = [3,3,3] or B = [4,4,4]
*/