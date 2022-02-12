func backspaceCompare(s string, t string) bool {
    ss := strings.Split(s, "")
    tt := strings.Split(t, "")
    last := max(len(ss), len(tt))
    
    var first, second string
    
    f_counter, s_counter := 0, 0
    
    for i := last - 1; i >= 0; i-- {
        if i <= len(ss) - 1 {
            if ss[i] == "#" {
                f_counter++
            } else if f_counter > 0 {
                f_counter--
            } else if f_counter == 0 {
                first += ss[i]
            }
        }
        if i <= len(tt) - 1 {
            if tt[i] == "#" {
                s_counter++
            } else if s_counter > 0 {
                s_counter--
            } else if s_counter == 0 {
                second += tt[i]
            }
        }
    }
    
    if first == second {
        return true
    }
    return false
    
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:
	Input: s = "ab#c", t = "ad#c"
	Output: true
	Explanation: Both s and t become "ac".

	Example 2:
	Input: s = "ab##", t = "c#d#"
	Output: true
	Explanation: Both s and t become "".

	Example 3:
	Input: s = "a##c", t = "#a#c"
	Output: true
	Explanation: Both s and t become "c".

	Example 4:
	Input: s = "a#c", t = "b"
	Output: false
	Explanation: s becomes "c" while t becomes "b".
*/