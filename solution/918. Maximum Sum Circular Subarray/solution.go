func maxSubarraySumCircular(A []int) (result int) {
    pos := false
    max_ending := -A[0]
    if len(A) == 1 {
        return A[0]
	}
	/* 
		Edge cases prevention
	*/
    if max_ending > 0 {
        max_ending = 0
        result = 0
    }else{
        result = -A[0]
    }
    all_sum := 0
    biggest := 0
    hmm := 0
    for _, value := range A {
        if value > 0 {
            pos = true
		}
		/* 
			Original array comparison
		*/
        hmm = max(value, hmm+value)
		biggest = max(biggest, hmm)
		/* 
			Inverted the array and compare
		*/
        value = -value
        all_sum += value
        max_ending = max(value, max_ending+value)
        result = max(result, max_ending)
    }
    if !pos {
		/* 
			Negative value comparison
		 */
        result = A[0]
        lol := A[0]
        for _, value := range A {
            lol = max(value, lol+value)
            result = max(result, lol)
        }
        return
    }
    return max(-(all_sum) - -(result), biggest)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:

	Input: [1,-2,3,-2]
	Output: 3
	Explanation: Subarray [3] has maximum sum 3

	Example 2:

	Input: [5,-3,5]
	Output: 10
	Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10

	Example 3:

	Input: [3,-1,2,-1]
	Output: 4
	Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4

	Example 4:

	Input: [3,-2,2,-3]
	Output: 3
	Explanation: Subarray [3] and [3,-2,2] both have maximum sum 3

	Example 5:

	Input: [-2,-3,-1]
	Output: -1
	Explanation: Subarray [-1] has maximum sum -1
*/