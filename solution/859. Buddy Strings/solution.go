func buddyStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// if a == b
	if a == b {
		mapping := make(map[string]bool)
		for _, value := range strings.Split(a, "") {
			if _, found := mapping[value]; found {
				return true
			}
			mapping[value] = true
		}
		return false
	}

	// if a != b
	var pairs [][]string
	aa, bb := strings.Split(a, ""), strings.Split(b, "")
	for i := 0; i < len(aa); i++ {
		if aa[i] != bb[i] {
			pairs = append(pairs, []string{aa[i], bb[i]})
		}
	}
	if len(pairs) >= 3 {
		return false
	}

	return len(pairs) == 2 && pairs[0][0] == pairs[1][1] && pairs[0][1] == pairs[1][0]
}

/*
	Example 1:
	Input: s = "ab", goal = "ba"
	Output: true
	Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.

	Example 2:
	Input: s = "ab", goal = "ab"
	Output: false
	Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.

	Example 3:
	Input: s = "aa", goal = "aa"
	Output: true
	Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.
*/