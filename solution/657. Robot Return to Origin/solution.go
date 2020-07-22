func judgeCircle(moves string) bool {
    m := strings.Split(moves, "")
    original := [2]int{0, 0}
    for _, value := range m {
        switch value {
            case "U": 
                original[0]++
                break
            case "R":
                original[1]++
                break
            case "D": 
                original[0]--
                break
            case "L":
                original[1]--
                break   
        }
    }
    fmt.Print(original)
    if original[0] == 0 && original[1] == 0 {
        return true
    }
    return false
}

/* 
Example 1:

Input: "UD"
Output: true 
Explanation: The robot moves up once, and then down once. All moves have the same magnitude, so it ended up at the origin where it started. Therefore, we return true.
 

Example 2:

Input: "LL"
Output: false
Explanation: The robot moves left twice. It ends up two "moves" to the left of the origin. We return false because it is not at the origin at the end of its moves.
*/