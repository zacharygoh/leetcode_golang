func diStringMatch(S string) (result []int) {
    start, end, last_appeared := 0, len(S), 0
    for _, value := range strings.Split(S, "") {
        switch value {
            case "I":
                result = append(result, start)
                start++
                last_appeared = start
                break
            case "D":
                result = append(result, end)
                end--
                last_appeared = end
                break
        }
    }
    result = append(result, last_appeared)
    return
}

/*
Example 1:

Input: "IDID"
Output: [0,4,1,3,2]

Example 2:

Input: "III"
Output: [0,1,2,3]

Example 3:

Input: "DDI"
Output: [3,2,0,1]
*/