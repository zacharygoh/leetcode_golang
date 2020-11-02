func minAddToMakeValid(S string) (result int) {
    left, right := 0, 0
    for _, value := range strings.Split(S, "") {
        if value == "(" {
            if right > 0 {
                result += right
                right = 0
            }
            left++
        }else{
            if left > 0 {
                left--
            }else{
                right++
            }
        }
    }
    return result + left + right
}

/* 
	Example 1:

	Input: "())"
	Output: 1

	Example 2:

	Input: "((("
	Output: 3

	Example 3:

	Input: "()"
	Output: 0

	Example 4:

	Input: "()))(("
	Output: 4
*/