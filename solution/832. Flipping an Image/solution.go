func flipAndInvertImage(A [][]int) [][]int {
    l := len(A[0])
    result := make([][]int, l)
    for i := 0; i < l; i++ {
        result[i] = make([]int, l)
    }
    for index, value := range A {
        for i, j := len(value)-1, 0; i > -1; i-- {
            result[index][i] = value[j]
            j++
        }
    }
    
    for _, value := range result {
        for index, v := range value {
            if v == 1 {
                value[index] = 0
            }else{
                value[index] = 1
            }
        }
    }
    return result
}

/* 
	Optimum solution
     for i := 0; i < len(A); i++ {
        low, high := 0, len(A)-1
        for low <= high {
            if A[i][low] == A[i][high] {
                A[i][low] = 1 - A[i][low]
                A[i][high] = A[i][low]
            }
            low++
            high--
        }
    }
    
    return A
*/

/* 
	Example 1:

	Input: [[1,1,0],[1,0,1],[0,0,0]]
	Output: [[1,0,0],[0,1,0],[1,1,1]]
	Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
	Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]

	Example 2:

	Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
	Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
	Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
	Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
*/