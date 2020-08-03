func countOdds(low int, high int) (result int) {
    for ; low <= high; low++ {
        if low % 2 != 0 {
            result++
        }
    }
    return
}

/* 
Example 1:
Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].

Example 2:
Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].
*/