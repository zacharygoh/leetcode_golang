func moveZeroes(nums []int)  {
    in := 0
    for i := 0; i < len(nums); i++ {
        if(nums[i] != 0) {
            nums[i], nums[in] = nums[in], nums[i]
            in++
        }
    }
}

/* 
Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
*/