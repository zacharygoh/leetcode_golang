func divisorGame(N int) bool {
    return N % 2 == 0
}

/* 
	Example 1:

	Input: 2
	Output: true
	Explanation: Alice chooses 1, and Bob has no more moves.

	Example 2:

	Input: 3
	Output: false
	Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
*/