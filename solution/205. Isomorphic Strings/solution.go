func isIsomorphic(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	isomorphic_map1 := make(map[byte]byte)
	isomorphic_map2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if value, found := isomorphic_map1[s[i]]; found && t[i] != value {
			return false
		}
		if value, found := isomorphic_map2[t[i]]; found && s[i] != value {
			return false
		}
		isomorphic_map1[s[i]] = t[i]
		isomorphic_map2[t[i]] = s[i]
	}
	return true
}

/* 
Example 1:

Input: s = "egg", t = "add"
Output: true

Example 2:

Input: s = "foo", t = "bar"
Output: false

Example 3:

Input: s = "paper", t = "title"
Output: true
*/