func runningSum(nums []int) []int {
    prev_num := 0
    for i := 0; i < len(nums); i++ {
        if(i+1 < len(nums)){
            prev_num = nums[i]
            nums[i+1] = prev_num + nums[i+1]
        }
    }
    return nums
}

/* 
Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

Example 2:

Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

Example 3:

Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]
*/