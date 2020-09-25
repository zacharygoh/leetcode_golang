func busyStudent(startTime []int, endTime []int, queryTime int) (result int) {
    for i := 0; i < len(startTime); i++ {
        if startTime[i] <= queryTime && queryTime <= endTime[i] {
            result++
        }
    }
    return result
}

/* 
	Example 1:

	Input: startTime = [1,2,3], endTime = [3,2,7], queryTime = 4
	Output: 1
	Explanation: We have 3 students where:
	The first student started doing homework at time 1 and finished at time 3 and wasn't doing anything at time 4.
	The second student started doing homework at time 2 and finished at time 2 and also wasn't doing anything at time 4.
	The third student started doing homework at time 3 and finished at time 7 and was the only student doing homework at time 4.

	Example 2:

	Input: startTime = [4], endTime = [4], queryTime = 4
	Output: 1
	Explanation: The only student was doing their homework at the queryTime.

	Example 3:

	Input: startTime = [4], endTime = [4], queryTime = 5
	Output: 0

	Example 4:

	Input: startTime = [1,1,1,1], endTime = [1,3,2,4], queryTime = 7
	Output: 0
	
	Example 5:

	Input: startTime = [9,8,7,6,5,4,3,2,1], endTime = [10,10,10,10,10,10,10,10,10], queryTime = 5
	Output: 5
*/