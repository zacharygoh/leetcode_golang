func stringMatching(words []string) (result []string) {
    wordss := words
    word_exist := make(map[string]bool)
    for index, value := range words {
        i := index
        i += 1
        for ; i < len(words); i++ {
            if strings.Contains(value, wordss[i]) {
                if _, found := word_exist[wordss[i]]; !found {
                    word_exist[wordss[i]] = true
                    result = append(result, wordss[i])   
                }
            }else if strings.Contains(wordss[i], value){
                if _, found := word_exist[value]; !found {
                    word_exist[value] = true
                    result = append(result, value)
                }
            }else{
                continue
            }
        }
    }
    return
}

/* 
Example 1:

Input: words = ["mass","as","hero","superhero"]
Output: ["as","hero"]
Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
["hero","as"] is also a valid answer.

Example 2:

Input: words = ["leetcode","et","code"]
Output: ["et","code"]
Explanation: "et", "code" are substring of "leetcode".

Example 3:

Input: words = ["blue","green","bu"]
Output: []
 
*/
