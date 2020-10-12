func restoreString(s string, indices []int) string {
    ss := strings.Split(s, "")
    c := make([]string, len(ss))
    for index, value := range indices {
        c[value] = ss[index]
    }
    return strings.Join(c, "")
}

/* 
	Example 1:

	Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
	Output: "leetcode"
	Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.

	Example 2:

	Input: s = "abc", indices = [0,1,2]
	Output: "abc"
	Explanation: After shuffling, each character remains in its position.

	Example 3:

	Input: s = "aiohn", indices = [3,1,4,2,0]
	Output: "nihao"

	Example 4:

	Input: s = "aaiougrt", indices = [4,0,2,6,7,3,1,5]
	Output: "arigatou"

	Example 5:

	Input: s = "art", indices = [1,0,2]
	Output: "rat"
*/