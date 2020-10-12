func relativeSortArray(arr1 []int, arr2 []int) (result []int) {
    maps := make(map[int]int)
    for _, value := range arr1 {
        maps[value]++
    }
    for _, value := range arr2 {
        result = append(result, value)
        if _, found := maps[value]; found {
            count := maps[value]
            for count > 1 {
                result = append(result, value)
                count--
            }
            delete(maps, value)
        }
    }
    mapp := make(map[int]bool)
    for _, value := range arr2 {
        mapp[value] = true
    }
    sort.Ints(arr1)
    for _, value := range arr1 {
        if _, found := mapp[value]; !found {
            result = append(result, value)
        }
    }
    return
}

/* 
	Example 1:

	Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
	Output: [2,2,2,1,4,3,3,9,6,7,19]
*/