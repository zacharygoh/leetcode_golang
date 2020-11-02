
func groupAnagrams(strs []string) (result [][]string) {
	mapping := make(map[string][]string)
    for _, value := range strs {
        s := sortString(value)
        mapping[s] = append(mapping[s], value)
    }
    for _, value := range mapping {
        result = append(result, value)
    }
    return 
/* 
	mapping := make(map[string]int)
    counter := 0
    for _, value := range strs {
        s := sortString(value)
        if _, found := mapping[s]; !found {
            mapping[s] = counter
            var a []string
            result = append(result, a)
            result[mapping[s]] = append(result[mapping[s]], value)
            counter++
        }else{
            result[mapping[s]] = append(result[mapping[s]], value)
        }
	} 
    return 
 */
}

func sortString(str string) string {
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

/* 
	Example 1:

	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	Example 2:

	Input: strs = [""]
	Output: [[""]]

	Example 3:

	Input: strs = ["a"]
	Output: [["a"]]
*/