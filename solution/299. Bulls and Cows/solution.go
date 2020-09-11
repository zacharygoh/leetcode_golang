func getHint(secret string, guess string) string {
    bull, cow := 0, 0
    cowHash := map[string]int{}
    
    sn, gn := strings.Split(secret, ""), strings.Split(guess, "")
    for i := 0; i < len(secret); i++ {
        if sn[i] != gn[i] {
            cowHash[sn[i]]++
        }
    }
    
    for i := 0; i < len(secret); i++ {
        if sn[i] == gn[i] {
            bull++
            continue
        }
        
        if value, found := cowHash[gn[i]]; found && value > 0 {
            cowHash[gn[i]]--
            cow++
        }
    }
    
    return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}

/* 
	Example 1:

	Input: secret = "1807", guess = "7810"

	Output: "1A3B"

	Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.

	Example 2:

	Input: secret = "1123", guess = "0111"

	Output: "1A1B"

	Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.
*/