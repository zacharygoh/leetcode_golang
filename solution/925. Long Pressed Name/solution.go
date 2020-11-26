func isLongPressedName(name string, typed string) bool {
    nn := strings.Split(name, "")
    tt := strings.Split(typed, "")
    name_length := len(nn)-1
    i := 0
    prev := ""
	
	/* fucking edge cases, should've done for first and last character */
    if nn[len(nn)-1] != tt[len(tt)-1] {
        return false        
    }
    
    for _, value := range tt {
        if nn[i] == value {
            if i == name_length {
                return true
            }
            i++
            prev = value
        }else{
            if prev != value {
                return false
            }
            prev = value
        }
    }
    return false
}

/* 
	Example 1:

	Input: name = "alex", typed = "aaleex"
	Output: true
	Explanation: 'a' and 'e' in 'alex' were long pressed.

	Example 2:

	Input: name = "saeed", typed = "ssaaedd"
	Output: false
	Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.

	Example 3:

	Input: name = "leelee", typed = "lleeelee"
	Output: true

	Example 4:

	Input: name = "laiden", typed = "laiden"
	Output: true
	Explanation: It's not necessary to long press any character.
*/