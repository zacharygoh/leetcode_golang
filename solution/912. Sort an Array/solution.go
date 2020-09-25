func sortArray(nums []int) []int {
    return quickSort(nums)
}

func quickSort(arr []int) []int {
    newArr := make([]int, len(arr))
    
    for index, value := range arr {
        newArr[index] = value
    }
    
    sort(newArr, 0, len(arr)-1)
    
    return newArr
}

func sort(arr []int, start, end int) {
    if end - start < 1 {
        return
    }
    
    pivot := arr[end]
    splitIndex := start
    
    for i := start; i < end; i++ {
        if arr[i] < pivot {
            temp := arr[splitIndex]
            
            arr[splitIndex], arr[i] = arr[i], temp
            
            splitIndex++
        }
    }
    
    arr[end], arr[splitIndex] = arr[splitIndex], pivot
    
    sort(arr, start, splitIndex-1)
    sort(arr, splitIndex+1, end)
}

/* 
	Example 1:

	Input: nums = [5,2,3,1]
	Output: [1,2,3,5]

	Example 2:

	Input: nums = [5,1,1,2,0,0]
	Output: [0,0,1,1,2,5]
*/