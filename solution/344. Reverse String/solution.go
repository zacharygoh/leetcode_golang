func reverseString(s []byte)  {
    first, last := 0, len(s)-1
    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}

/* 
	Example 1:

	Input: ["h","e","l","l","o"]
	Output: ["o","l","l","e","h"]

	Example 2:

	Input: ["H","a","n","n","a","h"]
	Output: ["h","a","n","n","a","H"]
*/