func shortestCompletingWord(licensePlate string, words []string) string {
    alphabet_map := make(map[rune]int)
    for _, value := range licensePlate {
        if unicode.IsLetter(value) {
            alphabet_map[unicode.ToLower(value)]++
        }
    }
    shortest := len(words[0])
    longest := 0
    for _, value := range words {
        if shortest > len(value) {
            shortest = len(value)
        }
        if longest < len(value) {
            longest = len(value)
        }
    }
    
    for i := shortest; i <= longest; i++ {
        for _, value := range words {
            if len(value) == i {
                if match(value, alphabet_map) {
                    return value
                }
            }
        }
    }
    return "-"
}

func match(word string, mapping map[rune]int) bool {
    alphabet_map := make(map[rune]int)
    for _, value := range word {
        alphabet_map[value]++
    }
    last_iteration := len(mapping) - 1
    i := 0
    for index, value := range mapping {
        if last_iteration == i {
            if alphabet_map[index] >= value {
                return true
            }else{
                return false
            }
        }
        if alphabet_map[index] >= value {
            i++
            continue
        }else{
            return false
        }
    }
    return false
}

type ByLen []string
func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

/* 
Example 1:
Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
Output: "steps"
Explanation: The smallest length word that contains the letters "S", "P", "S", and "T".
Note that the answer is not "step", because the letter "s" must occur in the word twice.
Also note that we ignored case for the purposes of comparing whether a letter exists in the word.

Example 2:
Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
Output: "pest"
Explanation: There are 3 smallest length words that contains the letters "s".
We return the one that occurred first.
*/