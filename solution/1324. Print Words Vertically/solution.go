func printVertically(s string) (result []string) {
	longest_word, index := 0, 0
    words := []string{}
    for _, value := range strings.Split(s, " ") {
		if len(value) > longest_word {
			longest_word = len(value)
		}
		words = append(words, value)
	}
	for i := 0; i < longest_word; i++ {
		var combined_letter string
		for _, value := range words {
			letters := strings.Split(value, "")
			if len(letters) > index {
				combined_letter+=letters[index]
			}else {
				combined_letter+=" "
			}
		}
        index++
		result = append(result, strings.TrimRight(combined_letter, " "))
	}
	return result
}

/* 
Example 1:

Input: s = "HOW ARE YOU"
Output: ["HAY","ORO","WEU"]
Explanation: Each word is printed vertically. 
 "HAY"
 "ORO"
 "WEU"

Example 2:

Input: s = "TO BE OR NOT TO BE"
Output: ["TBONTB","OEROOE","   T"]
Explanation: Trailing spaces is not allowed. 
"TBONTB"
"OEROOE"
"   T"

Example 3:

Input: s = "CONTEST IS COMING"
Output: ["CIC","OSO","N M","T I","E N","S G","T"]
*/