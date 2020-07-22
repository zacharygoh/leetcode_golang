func uncommonFromSentences(A string, B string) []string {
	words := make(map[string]int)
	for _, value := range strings.Split(A, " ") {
		words[value]++
	}
	for _, value := range strings.Split(B, " ") {
		words[value]++
	}
	unique_words := make([]string, 0, 200)
	for key, value := range words {
		if value == 1 {
			unique_words = append(unique_words, key)
		}
	}
	return unique_words
}

/* 
Example 1:

Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]

Example 2:

Input: A = "apple apple", B = "banana"
Output: ["banana"]
*/