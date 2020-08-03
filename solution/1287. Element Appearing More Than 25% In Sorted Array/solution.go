func findSpecialInteger(arr []int) (result int) {
    counter_map := make(map[int]int)
    most_appeared := 0
    for _, value := range arr {
        counter_map[value]++
        if counter_map[value] > most_appeared {
            most_appeared = counter_map[value]
            result = value
        }
    }
    return
}

/* 
Example 1:

Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6
*/