func isSubsequence(s string, t string) bool {
    texts := strings.Split(t, "")
    subsequent := strings.Split(s, "")
    leftover := 0
    if(len(subsequent) == 0){
        return true
    }
    for i := 0; i < len(texts); i++ {
        if(texts[i] == subsequent[leftover]){
            leftover++
            if(leftover == len(subsequent)){
                return true
            }
        }
    }
    return false
}

/* 
Example 1:

Input: s = "abc", t = "ahbgdc"
Output: tru
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

*/