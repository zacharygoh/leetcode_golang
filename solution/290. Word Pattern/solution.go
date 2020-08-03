func wordPattern(pattern string, str string) bool {
    strs := strings.Split(str, " ")
    string_map := make(map[byte]string)
    distinct_words := make(map[string]bool)
    
    if len(pattern) != len(strs) {
        return false
    }
    
    for i := 0; i < len(pattern); i++ {
        word, found := string_map[pattern[i]]
        
        if !found {
            if distinct := distinct_words[strs[i]]; distinct {
                return false
            }
            string_map[pattern[i]] = strs[i]
            distinct_words[strs[i]] = true
            continue
        }
        if strs[i] != word {
            return false
        }
    }
    return true
}

/* 
Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true

Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false

Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false

Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false
*/
