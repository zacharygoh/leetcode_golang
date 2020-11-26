func smallestRangeII(A []int, K int) (result int) {
    sort.Ints(A)
    last := len(A)-1
    result = A[last] - A[0]
    
    for i := 0; i < last; i++ {
        min := Min(A[0]+K, A[i+1]-K)
        max := Max(A[last]-K, A[i]+K)
        
        result = Min(result, max-min)
    }
    return
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
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
	Output: 3
	Explanation: B = [4,6,3]
*/