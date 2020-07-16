func findDisappearedNumbers(nums []int) []int {
    num_map := make(map[int]int)
    for _, value := range nums {
        num_map[value]++
    }
    var list []int
    for i := 1; i <= len(nums); i++ {
        if _, found := num_map[i]; !found {
            list = append(list, i)
        }
    }
    return list
}

/* 
Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
*/