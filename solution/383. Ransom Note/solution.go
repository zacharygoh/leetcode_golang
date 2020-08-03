func canConstruct(ransomNote string, magazine string) bool {
    ransom_map, magazine_map := make(map[string]int), make(map[string]int)
    
    rsn, mag := strings.Split(ransomNote, ""), strings.Split(magazine, "")
    
    length := longestLength(len(rsn), len(mag))
    
    for i := 0; i < length; i++ {
        if i < len(rsn) {
            ransom_map[rsn[i]]++
        }
        
        if i < len(mag) {
            magazine_map[mag[i]]++
        }
    }
    total := 0
    for key, value := range ransom_map {
        if _, found := magazine_map[key]; found {
            if magazine_map[key] >= value {
                total++
                if total == len(ransom_map) {
                    return true
                }
            }else{
                return false
            }
        }else{
            return false
        }
    }
    return true
}

func longestLength(first_array, second_array int) (result int) {
    if first_array > second_array {
        result = first_array
    }else{
        result = second_array
    }
    return
}

/* 
Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/