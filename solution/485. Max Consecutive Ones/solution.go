func findMaxConsecutiveOnes(nums []int) (result int) {
    counter := 0 
    for _, value := range nums {
        if value == 1 {
            counter++
            if result < counter {
                result = counter
            }
        }else{
            counter = 0 
        }
    }
    return
}

/* 
	Example 1:
	Input: [1,1,0,1,1,1]
	Output: 3
	Explanation: The first two digits or the last three digits are consecutive 1s.
		The maximum number of consecutive 1s is 3.
*/