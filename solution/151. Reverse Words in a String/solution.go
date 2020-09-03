func reverseWords(s string) (result string) {
    words := strings.Split(s, " ")
    for _, value := range words {
        value = strings.TrimRight(value, " ")
        value = strings.TrimLeft(value, " ")
        if value != "" {
            result = value + " " + result
        }
    }
    result = strings.TrimRight(result, " ")
    return 
}

/* 
	Example 1:

	Input: "the sky is blue"
	Output: "blue is sky the"

	Example 2:

	Input: "  hello world!  "
	Output: "world! hello"
	Explanation: Your reversed string should not contain leading or trailing spaces.

	Example 3:

	Input: "a good   example"
	Output: "example good a"
	Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/