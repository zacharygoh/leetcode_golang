func dominantIndex(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    temp_table := make(map[int]int)
    for index, value := range nums {
        if _, found := temp_table[value]; found {
            continue
        }
        temp_table[value] = index
    }
    sort.Ints(nums)
    largest_num := nums[len(nums)-1]
    for i := len(nums)-2; i > -1; i-- {
        if nums[i] * 2 <= largest_num {
            return temp_table[largest_num]
        } else if largest_num == nums[i] {
            continue   
        }else{
            return -1
        }
    }
    return -1
}

/* 
Example 1:

Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
 

Example 2:

Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
*/