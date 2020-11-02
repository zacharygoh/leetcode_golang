func canThreePartsEqualSum(A []int) bool {
    total := 0
    for _, value := range A {
        total += value
    }
    expectedSum, count, sum := total/3, 0, 0
    for _, value := range A {
        sum += value
        if sum == expectedSum {
            sum = 0
            count++
        }
        if count == 3 {
            return true
        }
    }
    return false
}

/* 
	Example 1:

	Input: A = [0,2,1,-6,6,-7,9,1,2,0,1]
	Output: true
	Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1

	Example 2:

	Input: A = [0,2,1,-6,6,7,9,-1,2,0,1]
	Output: false

	Example 3:

	Input: A = [3,3,6,5,-2,2,5,1,-9,4]
	Output: true
	Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
*/