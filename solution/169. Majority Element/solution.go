func majorityElement(nums []int) int {
    appear := 0
    output := 0
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if _, found := m[nums[i]]; found {
            m[nums[i]]++
        }else{
            m[nums[i]] = 1
        }
    }
    for key, value := range m {
        if(value >= appear){
            appear = value
            output = key
        }
    }
    return output
}

/* 
    Example 1:

    Input: [3,2,3]
    Output: 3

    Example 2:

    Input: [2,2,1,1,1,2,2]
    Output: 2
*/