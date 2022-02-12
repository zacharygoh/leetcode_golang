func decodeString(s string) string {
    return decode(s, 1)
}

func decode(s string, n int) (result string) {
    start, brac, times := 0, 0, 0
    flag := true
    for i, c := range s {
        if '0' <= c && c <= '9' {
            if flag {
                times = times * 10 + int(c - '0')
            }
        } else if c == '[' {
            if brac == 0 {
                start = i
            }
            brac++
            flag = false
        } else if c == ']' {
            brac--
            if brac == 0 {
                result += decode(s[start+1:i], times)
                times = 0
                flag = true
            }
        } else {
            if brac == 0 {
                result = result + string(c)
            }
        }
    }
    temp := ""
    for n > 0 {
        temp += result
        n--
    }
    return temp
}

/* 
	Example 1:
	Input: s = "3[a]2[bc]"
	Output: "aaabcbc"

	Example 2:
	Input: s = "3[a2[c]]"
	Output: "accaccacc"

	Example 3:
	Input: s = "2[abc]3[cd]ef"
	Output: "abcabccdcdcdef"

	Example 4:
	Input: s = "abc3[cd]xyz"
	Output: "abccdcdcdxyz"
*/