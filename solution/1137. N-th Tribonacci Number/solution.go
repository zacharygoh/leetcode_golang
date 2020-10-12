func tribonacci(n int) int {
    if n < 2 {
        return n
    }
    trib, prev, prev_1 := 1, 1, 0
    for n > 2 {
        trib, prev, prev_1 = trib + prev + prev_1, trib, prev
        n--
    }
    return trib
}

/* 
	Example 1:

	Input: n = 4
	Output: 4
	Explanation:
	T_3 = 0 + 1 + 1 = 2
	T_4 = 1 + 1 + 2 = 4

	Example 2:

	Input: n = 25
	Output: 1389537
*/