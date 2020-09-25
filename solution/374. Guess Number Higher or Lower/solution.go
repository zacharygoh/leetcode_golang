/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

 func guessNumber(n int) (result int) {
    if guess(n) == -1 {
        result = 1
    }
    if guess(n) == 1 {
        result = 2
    }
    if guess(n) == 0 {
        return n
    }
    if result == 1 {
        for i := n; i > -1; i-- {
            if guess(i) == 0 {
                return i
            }
        }
    }
    if result == 2 {
        return n
        // for i := 0; i < n; i++ {
        //     if guess(i) == 0 {
        //         return i
        //     }
        // }
    }
    return
}

/* 
	-1 : My number is lower
	1 : My number is higher
	0 : Congrats! You got it!
	Example :

	Input: n = 10, pick = 6
	Output: 6
*/