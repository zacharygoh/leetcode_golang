/* 
	O(nlogn+n)
*/
func minIncrementForUnique(A []int) (result int) {
    sort.Ints(A)
    for i := 1; i < len(A); i++ {
        prev, curr := A[i - 1], A[i]
        if prev >= curr {
            result += prev - curr + 1
            A[i] = prev + 1
        }
    }
    return
}

/* 
	O(mn)
*/
// func minIncrementForUnique(A []int) (result int) {
//     if len(A) < 2 {
//         return 
//     }
    
//     count := make([]int, len(A)*2)
//     for _, value := range A {
//         count[value]++
//     }
    
//     for index, value := range count {
//         if value > 1 {
//             result += value - 1
//             count[index + 1] += value - 1
//         }
//     }
//     return
// }

/* 
	Example 1:

	Input: [1,2,2]
	Output: 1
	Explanation:  After 1 move, the array could be [1, 2, 3].

	Example 2:

	Input: [3,2,1,2,1,7]
	Output: 6
	Explanation:  After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
	It can be shown with 5 or less moves that it is impossible for the array to have all unique values.
*/