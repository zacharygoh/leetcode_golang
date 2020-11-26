func longestOnes(A []int, K int) int {
    /* Catch up solution */
    i, j := 0, 0
    for ; i < len(A); i++ {
        K -= 1 - A[i]
        if K < 0 {
            K += 1 - A[j]
            j++
        }
    }
    return i - j
}

/* 
	Example 1:

	Input: A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
	Output: 6
	Explanation: 
	[1,1,1,0,0,1,1,1,1,1,1]
	Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.

	Example 2:

	Input: A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
	Output: 10
	Explanation: 
	[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
	Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.
*/