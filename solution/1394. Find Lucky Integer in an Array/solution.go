func findLucky(arr []int) int {
    lucky_map := make(map[int]int)
    for _, value := range arr {
        if _, found := lucky_map[value]; found {
            lucky_map[value]++
        }else{
            lucky_map[value] = 1
        }
    }
    appeared := 0
    for key, value := range lucky_map {
        if(key == value && appeared < value){
            appeared = value
        } 
    }
    if(appeared != 0){
        return appeared
    }else{
        return -1
    }
}

/* 
Example 1:

Input: arr = [2,2,3,4]
Output: 2
Explanation: The only lucky number in the array is 2 because frequency[2] == 2.

Example 2:

Input: arr = [1,2,2,3,3,3]
Output: 3
Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.

Example 3:

Input: arr = [2,2,2,3,3]
Output: -1
Explanation: There are no lucky numbers in the array.

Example 4:

Input: arr = [5]
Output: -1

Example 5:

Input: arr = [7,7,7,7,7,7,7]
Output: 7
*/