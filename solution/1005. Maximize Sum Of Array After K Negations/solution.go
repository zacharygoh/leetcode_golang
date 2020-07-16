func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	total_negative := 0
	has_negative := false
	for i := 0; i < K; i++ {
		if len(A)-1 < i {
			break
		}
		if A[i] < 0 {
			A[i]*=-1
			total_negative = i
			has_negative = true
		}
	}
    sort.Ints(A)
	if (K-total_negative % 2) % 2 == 0 || (K % 2 > 0 && !has_negative) {
		A[0]*=-1
	}
	result := 0
	for _, value := range A {
		result+=value
	}
	return result
}

/* 
Example 1:

Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].

Example 2:

Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].

Example 3:

Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
 */