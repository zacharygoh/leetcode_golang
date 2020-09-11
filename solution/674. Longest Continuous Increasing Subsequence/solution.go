func findLengthOfLCIS(nums []int) (counter int) {
    latest := 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] < nums[i] {
            latest++
            if counter < latest {
                counter = latest
            }
        }else{
            latest = 0
        }
    }
    if len(nums) == 0 {
        counter = 0
    }else{
        counter++
    }
    return
}

/* 
Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3. 
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4. 

Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1. 
*/