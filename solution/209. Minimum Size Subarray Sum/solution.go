func minSubArrayLen(s int, nums []int) (result int) {
    left, sum := 0, 0
    result = 10000000
    for i, value := range nums {
        sum += value
        for sum >= s {
            result = min(result, i+1-left)
            sum -= nums[left]
            left++
        }
    }
    if result != 10000000 {
        return
    }else{
        return 0
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

/* 
	Example: 

	Input: s = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: the subarray [4,3] has the minimal length under the problem constraint.
*/