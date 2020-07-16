func mostCommonWord(paragraph string, banned []string) string {
    word_map := make(map[string]int)
    paragraph = strings.ToLower(paragraph)
    words := splitAny(paragraph, " !?,.';")
    for i := 0; i < len(words); i++ {
        if _, found := word_map[words[i]]; found {
            word_map[words[i]]++
        }else{ 
            word_map[words[i]] = 1
        }
    }
    appeared := 0
    var result string
    for key, value := range word_map {
        if(contains(banned, key) == false){
            if(appeared < value){
                appeared = value
                result = key
            }
        }
    }
    return result
}

func splitAny(text string, seperator string) []string {
    splitter := func(r rune) bool {
        return strings.ContainsRune(seperator, r)
    }
    return strings.FieldsFunc(text, splitter)
}

func contains(arr []string, matched string) bool {
    for _, value := range arr {
        if value == matched {
            return true
        }
    }
    return false
}

/* 
Example:

Input: 
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation: 
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"), 
and that "hit" isn't the answer even though it occurs more because it is banned.
*/