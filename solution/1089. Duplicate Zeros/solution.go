func duplicateZeros(arr []int)  {
    for i := 0; i < len(arr) - 1; i++ {
        if arr[i] == 0 {
            for j := len(arr) - 1; j > i; j-- {
                arr[j] = arr[j - 1]
            }
            i++
        }
    }
}

/* 
	Example 1:

	Input: [1,0,2,3,0,4,5,0]
	Output: null
	Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

	
	Example 2:

	Input: [1,2,3]
	Output: null
	Explanation: After calling your function, the input array is modified to: [1,2,3]
*/