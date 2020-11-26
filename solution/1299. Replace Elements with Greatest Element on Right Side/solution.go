func replaceElements(arr []int) []int {
    last := len(arr)-1
    mx := -1
    for i := last; i > -1; i-- {
        arr[i], mx = mx, max(arr[i], mx)
    }
    return arr
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:

	Input: arr = [17,18,5,4,6,1]
	Output: [18,6,6,6,1,-1]
*/