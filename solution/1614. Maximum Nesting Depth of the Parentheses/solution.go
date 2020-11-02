func maxDepth(s string) (result int) {
    max := 0
    for _, value := range strings.Split(s, "") {
        if value == "(" {
            max++
        }
        if value == ")" {
            max--
        }
        if result < max {
            result = max
        }
    }
    return 
}

/* 
	Example 1:

	Input: s = "(1+(2*3)+((8)/4))+1"
	Output: 3
	Explanation: Digit 8 is inside of 3 nested parentheses in the string.

	Example 2:

	Input: s = "(1)+((2))+(((3)))"
	Output: 3

	Example 3:

	Input: s = "1+(2*3)/(2-1)"
	Output: 1

	Example 4:

	Input: s = "1"
	Output: 0
*/