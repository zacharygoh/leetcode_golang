func maxNumberOfBalloons(text string) (result int) {
    letter_map := make(map[string]int)
    for _, value := range strings.Split(text, "") {
        letter_map[value]++
    }
    //balloon
    var b = []string{"b", "a", "l", "o", "n"}
    for i := 0; i < len(text) / 7; i++ {
        for j := 0; j < 5; j++ {
            switch b[j] {
                case "b":
                    if letter_map[b[j]] >= 1 {
                        letter_map[b[j]]--
                    }else{
                        return
                    }
                    //1
                    break
                case "a":
                    if letter_map[b[j]] >= 1 {
                        letter_map[b[j]]--
                    }else{
                        return
                    }
                    //1
                    break
                case "l":
                    if letter_map[b[j]] >= 2 {
                        letter_map[b[j]]-=2
                    }else{
                        return
                    }
                    //2
                    break
                case "o":
                    if letter_map[b[j]] >= 2 {
                        letter_map[b[j]]-=2
                    }else{
                        return
                    }
                    //2
                    break
                case "n":
                    if letter_map[b[j]] >= 1 {
                        letter_map[b[j]]--
                    }else{
                        return
                    }
                    //1
                    break
            }
        }
        result++
    }
    return
}

/* 
	Example 1:

	Input: text = "nlaebolko"
	Output: 1

	Example 2:

	Input: text = "loonbalxballpoon"
	Output: 2

	Example 3:

	Input: text = "leetcode"
	Output: 0
*/