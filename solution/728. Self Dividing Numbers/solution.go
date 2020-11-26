func selfDividingNumbers(left int, right int) (result []int) {
    for left <= right {
        temp := left
        for temp > 0 {
            digit := temp % 10
            if !(digit != 0 && left % digit == 0) {
                break
            }
            temp /= 10
        }
        if temp == 0 {
            result = append(result, left)
        }
        left++
    }
    return
}

/* 
	Example 1:
	Input: 
	left = 1, right = 22
	Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
*/