func reverseOnlyLetters(S string) string {
	ss := strings.Split(S, "")
	result := ""
	j := len(ss)-1
	for i := 0; i < len(ss); i++ {
		if !IsLetter(ss[i]) {
			result+=ss[i]
            continue
		}
		
		for ; j > -1; j-- {
			if !IsLetter(ss[j]) {
				continue
			}
            result+=ss[j]
            j--
            break
		}
	}
	return result
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/* 
Example 1:

Input: "ab-cd"
Output: "dc-ba"

Example 2:

Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"

Example 3:

Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"
*/