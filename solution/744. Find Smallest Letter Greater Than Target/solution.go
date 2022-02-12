package main

func nextGreatestLetter(letters []byte, target byte) (result byte) {
	for _, value := range letters {
		if target < value {
			result = min(value, target)
			break
		}
	}
	if result == 0 {
		result = letters[0]
	}
	return
}

func min(a, b byte) byte {
	if a < b {
		return b
	}
	return a
}

/*
	Assumption because in non-decreasing order = increasing order

	Example 1:
	Input: letters = ["c","f","j"], target = "a"
	Output: "c"

	Example 2:
	Input: letters = ["c","f","j"], target = "c"
	Output: "f"

	Example 3:
	Input: letters = ["c","f","j"], target = "d"
	Output: "f"
*/
