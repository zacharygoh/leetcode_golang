func singleNumber(nums []int) int {
    appeared_map := make(map[int]int)
    for _, value := range nums {
        appeared_map[value]++
    }
    for key, value := range appeared_map {
        if value == 1 {
            return key
        }
    }
    return 0
}
/* 
Example 1:

Input: [2,2,3,2]
Output: 3

Example 2:

Input: [0,1,0,1,0,1,99]
Output: 99
*/