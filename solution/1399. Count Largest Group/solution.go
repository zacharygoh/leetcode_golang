func countLargestGroup(n int) int {
    counter_map := make(map[string]int)
    combine_value, largest_group := 0, 0
    
    for i := n; i > -1; i-- {
        value := strconv.Itoa(i)
        if len(value) > 1 {
            for _, value1 := range strings.Split(value, "") {
                test, _ := strconv.Atoi(value1)
                combine_value+=test
            }
            counter_map[strconv.Itoa(combine_value)]++
            combine_value = 0
            continue
        }
        counter_map[value]++
    }
    
    combine_value = 0
    for key, value := range counter_map {
        if key != "0" {
            if value > largest_group {
                combine_value = 1
                largest_group = value
                continue
            }else if value == largest_group {
                combine_value++
                continue
            }
        }
    }
    return combine_value
}

/* 
Example 1:

Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.

Example 2:

Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.

Example 3:

Input: n = 15
Output: 6

Example 4:

Input: n = 24
Output: 5
*/