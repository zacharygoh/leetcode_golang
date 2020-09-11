func reverseWords(s string) (result string) {
    ss := strings.Split(s, " ")
    last_index := len(ss)-1
    for index, value := range ss {
        sss := strings.Split(value, "")
        li := len(sss)-1
        for j := li; j > -1; j-- {
            result += sss[j]
        }
        if index != last_index {
            result += " "
        }
    }
    return
}

/* 
	Example 1:
	Input: "Let's take LeetCode contest"
	Output: "s'teL ekat edoCteeL tsetnoc"
*/