func countCharacters(words []string, chars string) (result int) {
    char_map := make(map[rune]int)
    for _, value := range chars {
        char_map[value]++
    }
    for _, value := range words {
        letter_map := make(map[rune]int)
        first_counter := 0
        for _, char := range value {
            first_counter++
            letter_map[char]++
        }
        counter := 0
        for index, char := range letter_map {
            if char_map[index] >= char {
                counter += char
            }
        }
        if counter == first_counter {
            result += counter
        }
    }
    return
}

/* 
	Example 1:

	Input: words = ["cat","bt","hat","tree"], chars = "atach"
	Output: 6
	Explanation: 
	The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.

	Example 2:

	Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
	Output: 10
	Explanation: 
	The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
*/