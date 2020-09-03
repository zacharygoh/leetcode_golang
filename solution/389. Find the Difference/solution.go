func findTheDifference(s string, t string) (result byte) {
    first_map, second_map := make(map[rune]int), make(map[rune]int)
    for _, value := range s {
        first_map[value]++
    }
    for _, value := range t {
        second_map[value]++
    }
    for key, value := range second_map {
        if first_map[key] != value {
            return byte(key)
        }
    }
    return
}

/* 
	Example:

	Input:
	s = "abcd"
	t = "abcde"

	Output:
	e

	Explanation:
	'e' is the letter that was added.
*/