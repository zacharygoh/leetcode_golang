func largeGroupPositions(S string) (result [][]int) {
    ss := strings.Split(S, "")
    i := 1
    for index, value := range ss {
        counter, end := 0, 0
        for ; i < len(ss); i++ {
            if value == ss[i] {
                end = i
                counter++
            }
            if value != ss[i] {
                if end - index >= 2 {
                   result = append(result, []int{index, end})
                }
                break
            }
            if end + 1 == len(ss) && end - index >= 2 {
                result = append(result, []int{index, end})
            }
        }
    }
    return
}

/*

Example 1:

Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.

Example 2:

Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.

Example 3:

Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
*/