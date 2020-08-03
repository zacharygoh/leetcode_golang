func getNoZeroIntegers(n int) (result []int) {
    for i := 1; i < n; i++ {
        remainder := n - i
        if !strings.Contains(strconv.Itoa(remainder), "0") && !strings.Contains(strconv.Itoa(i), "0") {
            result = append(result, i)
            result = append(result, remainder)
            return
        }
    }
    return
}

/* 
Example 1:

Input: n = 2
Output: [1,1]
Explanation: A = 1, B = 1. A + B = n and both A and B don't contain any 0 in their decimal representation.
Example 2:

Input: n = 11
Output: [2,9]
Example 3:

Input: n = 10000
Output: [1,9999]
Example 4:

Input: n = 69
Output: [1,68]
Example 5:

Input: n = 1010
Output: [11,999]
*/
