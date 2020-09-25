func isPowerOfThree(n int) bool {
    if n < 1 {
        return false
    }

	/* O(n) */
	//     for n % 3 == 0 {
	//         n /= 3
	//     }
		
	//     return n == 1

	/* O(1) */
    return 1162261467 % n == 0
}

/* 
	Example 1:

	Input: 27
	Output: true

	Example 2:

	Input: 0
	Output: false

	Example 3:

	Input: 9
	Output: true

	Example 4:

	Input: 45
	Output: false
*/