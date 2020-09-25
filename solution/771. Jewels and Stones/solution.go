func numJewelsInStones(J string, S string) (result int) {
    letter_map := make(map[rune]bool)
    for _, value := range J {
        letter_map[value] = true
    }
    
    for _, value := range S {
        if letter_map[value] {
            result++
        }
    }
    return 
}

/* 
	Example 1:

	Input: J = "aA", S = "aAAbbbb"
	Output: 3
	
	Example 2:

	Input: J = "z", S = "ZZ"
	Output: 0
*/