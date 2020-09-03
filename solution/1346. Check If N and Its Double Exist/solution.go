func checkIfExist(arr []int) bool {
    sort.Ints(arr)
    most_right := len(arr)-1
    counter := 0
    for index, _ := range arr {
        i := len(arr) - (index + 2)
        for ; i > -1; i-- {
            if arr[i] == 0 {
                counter++
                if len(arr)-1 == counter {
                    return true
                }
                continue
            }
            if float64(arr[most_right]) / float64(arr[i]) == 2 {
                return true
            }
            if float64(arr[i]) / float64(arr[most_right]) == 2 {
                return true
            }
        }
        most_right--
    }
    return false
}

/* 
	Example 1:

	Input: arr = [10,2,5,3]
	Output: true
	Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.

	Example 2:

	Input: arr = [7,1,14,11]
	Output: true
	Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
	
	Example 3:

	Input: arr = [3,1,7,11]
	Output: false
	Explanation: In this case does not exist N and M, such that N = 2 * M.
*/