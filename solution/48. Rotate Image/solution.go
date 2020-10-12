func rotate(matrix [][]int)  {
    length := len(matrix)
    for i := 0; i < length; i++ {
        for j := i; j < length; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    for length > 0 {
        length--
        for i, j := 0, len(matrix)-1; i < len(matrix)/2; i++ {
            matrix[length][i], matrix[length][j] = matrix[length][j], matrix[length][i]
            j--
        }
    }
}

/* 
	Example 1:

	Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
	Output: [[7,4,1],[8,5,2],[9,6,3]]

	Example 2:

	Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
	Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

	Example 3:

	Input: matrix = [[1]]
	Output: [[1]]

	Example 4:

	Input: matrix = [[1,2],[3,4]]
	Output: [[3,1],[4,2]]
*/