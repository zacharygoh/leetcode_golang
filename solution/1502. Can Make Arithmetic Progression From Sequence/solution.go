func canMakeArithmeticProgression(arr []int) bool {
    quickSort(arr)
    differences := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        if arr[i] - arr[i-1] != differences {
            return false
        }
    }
    return true
}

/* Quick sort */
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

	Input: arr = [3,5,1]
	Output: true
	Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.

	Example 2:

	Input: arr = [1,2,4]
	Output: false
	Explanation: There is no way to reorder the elements to obtain an arithmetic progression.
*/