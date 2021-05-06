/* 
	Brute force solution
 */
func dailyTemperatures(T []int) (result []int) {
    top, top_i := 0, 0
    for index, value := range T {
        temp := 0
        for i := index + 1; i < len(T); i++ {
            if value < T[i] {
                temp = i - index
                break
            }
        }
        result = append(result, temp)
    }
    return
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

/* 
	Monotonic stack approach
*/

func dailyTemperatures(T []int) []int {
    stack := []int{}
    result := make([]int, len(T))
    for i := 0; i < len(T); i++ {
        for len(stack) != 0 && T[stack[len(stack) - 1]] < T[i] {
            result[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }
    return result
}

/* 
	For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
*/