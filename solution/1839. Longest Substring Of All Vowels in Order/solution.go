func longestBeautifulSubstring(word string) (counter int) {
    words := strings.Split(word, "")
    temp_counter := 1
    c := 0
    
    if len(words) < 5 {
        return
    }
    
    
    for i := 0; i < len(words); i++ {
        ii := i + 1
        if i == len(words) - 1 {
            ii = i
        }
        if words[i] == "a" && temp_counter == 1 {
            c++
            if words[ii] == "e" {
                temp_counter = 2
            } else if words[ii] != "a" {
                temp_counter = 1
                c = 0
            }
        } else if words[i] == "e" && temp_counter == 2 {
            c++
            if words[ii] == "i" {
                temp_counter = 3
            } else if words[ii] != "e" {
                temp_counter = 1
                c = 0
            }
        } else if words[i] == "i" && temp_counter == 3 {
            c++
            if words[ii] == "o" {
                temp_counter = 4
            } else if words[ii] != "i" {
                temp_counter = 1
                c = 0
            }
        } else if words[i] == "o" && temp_counter == 4 {
            c++
            if words[ii] == "u" {
                temp_counter = 5
            } else if words[ii] != "o" {
                temp_counter = 1
                c = 0
            }
        } else if words[i] == "u" && temp_counter == 5 {
            c++
            counter = max(counter, c)
            if words[ii] != "u" {
                temp_counter = 1
                c = 0
            }
        }
    }
    return 
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:
	Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
	Output: 13
	Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.
	
	Example 2:
	Input: word = "aeeeiiiioooauuuaeiou"
	Output: 5
	Explanation: The longest beautiful substring in word is "aeiou" of length 5.
	
	Example 3:
	Input: word = "a"
	Output: 0
	Explanation: There is no beautiful substring, so return 0.
*/