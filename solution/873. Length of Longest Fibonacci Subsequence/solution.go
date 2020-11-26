func lenLongestFibSubseq(A []int) (result int) {
    mapping := make(map[int]bool)
    for _, value := range A {
        mapping[value] = true
    }
    for i := range A {
        for j := i + 1; j < len(A); j++ {
            a, b := A[i], A[j]
            c := a + b
            l := 2
            for mapping[c] {
                a = b
                b = c
                c = a + b
                l++
                result = max(result, l)
            }
        }
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:

	Input: [1,2,3,4,5,6,7,8]
	Output: 5
	Explanation:
	The longest subsequence that is fibonacci-like: [1,2,3,5,8].

	Example 2:

	Input: [1,3,7,11,12,14,18]
	Output: 3
	Explanation:
	The longest subsequence that is fibonacci-like:
	[1,11,12], [3,11,14] or [7,11,18].
*/