func canWinNim(n int) bool {
    if n % 4 != 0 {
        return true
    }   
    return false
}

/* 
Example:

Input: 4
Output: false 
Explanation: If there are 4 stones in the heap, then you will never win the game;
             No matter 1, 2, or 3 stones you remove, the last stone will always be 
             removed by your friend.
*/