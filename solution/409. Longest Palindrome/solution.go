func longestPalindrome(s string) int {
	letter_map := make(map[string]int)
	for _, value := range strings.Split(s, "") {
		letter_map[value]++
	}
	result, missing_piece := 0, 1
	for _, value := range letter_map {
		if value % 2 == 0 {
			result+=value
		}else{
			result+=(value-1)
		}
	}
	return result + missing_piece
}

/* 
Example:

Input:
"abccccdd"

Output:
7
*/