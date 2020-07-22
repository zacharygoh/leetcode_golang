func uniqueOccurrences(arr []int) bool {
    counter_map := make(map[int]int)
    for _, value := range arr {
        counter_map[value]++
    }
    occurance := make(map[int]int)
    for _, value := range counter_map {
        if _, found := occurance[value]; !found {
            occurance[value] = value
        }else{
            return false
        }
    }
    return true
}

/* 
Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

Example 2:

Input: arr = [1,2]
Output: false

Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true
*/