func maxPower(s string) (result int) {
    appeared := 1
    result = 1
    for i := 1; i < len(s); i++ {
        if s[i - 1] == s[i] {
            appeared++
        }else{
            appeared = 1
        }
        if appeared > result {
            result = appeared
        }
    }
    return
}

/* 
Example 1:

Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.

Example 2:

Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.

Example 3:

Input: s = "triplepillooooow"
Output: 5

Example 4:

Input: s = "hooraaaaaaaaaaay"
Output: 11
Example 5:

Input: s = "tourist"
Output: 1
*/