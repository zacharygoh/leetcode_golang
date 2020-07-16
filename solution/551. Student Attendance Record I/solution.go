func checkRecord(s string) bool {
	ss := strings.Split(s, "")
	previous_value := ""
	a_counter, l_counter := 0, 1
	for i := 0; i < len(ss); i++ {
		if ss[i] != "L" {
			previous_value = ""
			l_counter = 1
		}
		switch (ss[i]) {
		case "L":
			if previous_value == ss[i] {
				l_counter++
				if l_counter > 2 {
					return false
				}
			}
			break
		case "A":
			a_counter++
			if a_counter > 1 {
				return false
			}
			break
		}
		previous_value = ss[i]
	}
	return true
}

/* 
Example 1:
Input: "PPALLP"
Output: True

Example 2:
Input: "PPALLL"
Output: False
*/