func peakIndexInMountainArray(arr []int) (result int) {
    for index, value := range arr {
        if value > arr[index+1] {
            result = index
            break
        }
    }
    return
}

/* 
	Example 1:

	Input: arr = [0,1,0]
	Output: 1

	Example 2:

	Input: arr = [0,2,1,0]
	Output: 1

	Example 3:

	Input: arr = [0,10,5,2]
	Output: 1

	Example 4:

	Input: arr = [3,4,5,1]
	Output: 2

	Example 5:

	Input: arr = [24,69,100,99,79,78,67,36,26,19]
	Output: 2
*/