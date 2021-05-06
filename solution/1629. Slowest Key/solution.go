func slowestKey(releaseTimes []int, keysPressed string) byte {
    mapping := make(map[int]int)
    start, max, i := 0, 0, 0
    var a byte
    for index, value := range releaseTimes {
        max = Max(max, value - start)
        mapping[max]++
        if max == value - start {
            i = index
        }
        start = value
    }
    if mapping[max] > 1 {
        start = 0
        for i, v := range releaseTimes {
            if max == v - start {
                a = LetterMax(keysPressed[i], a)
            }
            start = v
        }
        return a
    }else{
        return keysPressed[i]
    }
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func LetterMax(a, b byte) byte {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:

	Input: releaseTimes = [9,29,49,50], keysPressed = "cbcd"
	Output: "c"
	Explanation: The keypresses were as follows:
	Keypress for 'c' had a duration of 9 (pressed at time 0 and released at time 9).
	Keypress for 'b' had a duration of 29 - 9 = 20 (pressed at time 9 right after the release of the previous character and released at time 29).
	Keypress for 'c' had a duration of 49 - 29 = 20 (pressed at time 29 right after the release of the previous character and released at time 49).
	Keypress for 'd' had a duration of 50 - 49 = 1 (pressed at time 49 right after the release of the previous character and released at time 50).
	The longest of these was the keypress for 'b' and the second keypress for 'c', both with duration 20.
	'c' is lexicographically larger than 'b', so the answer is 'c'.

	Example 2:

	Input: releaseTimes = [12,23,36,46,62], keysPressed = "spuda"
	Output: "a"
	Explanation: The keypresses were as follows:
	Keypress for 's' had a duration of 12.
	Keypress for 'p' had a duration of 23 - 12 = 11.
	Keypress for 'u' had a duration of 36 - 23 = 13.
	Keypress for 'd' had a duration of 46 - 36 = 10.
	Keypress for 'a' had a duration of 62 - 46 = 16.
	The longest of these was the keypress for 'a' with duration 16.
*/