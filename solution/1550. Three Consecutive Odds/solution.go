func threeConsecutiveOdds(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    for i, _ := range arr {
        if i + 2 < len(arr) {
            if arr[i] % 2 != 0 && arr[i+1] % 2 != 0 && arr[i+2] % 2 != 0 {
                return true
            }
        }
    }
    return false
    // counter := 0
    // for _, value := range arr {
    //     if value % 2 != 0 {
    //         counter++
    //         if counter == 3 {
    //             return true
    //         }
    //     }else{
    //         counter = 0
    //     }
    // }
    // return false
}

/* 
Example 1:

Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.

Example 2:

Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
Explanation: [5,7,23] are three consecutive odds.
*/