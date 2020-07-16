func repeatedNTimes(A []int) int {
    number_map := make(map[int]int)
    for _, value := range A {
        if number_map[value] > 0 {
            return value
        }
        number_map[value]++
    }
    return 0
}

/* 
Example 1:

Input: [1,2,3,3]
Output: 3

Example 2:

Input: [2,1,2,5,3,2]
Output: 2

Example 3:

Input: [5,1,5,2,5,3,5,4]
Output: 5
*/