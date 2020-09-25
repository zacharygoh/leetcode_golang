func isValidSudoku(board [][]byte) bool {
    valid_map := make(map[string]bool)
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            board_coordinate := string(board[i][j])
            if board_coordinate != "." {
                row := "Row " + strconv.Itoa(i) + " " + board_coordinate + " exists"
                col := "Column " + strconv.Itoa(j) + " " + board_coordinate + " exists"
                area := "Area " + strconv.Itoa((i/3)*3 + j/3) + " " + board_coordinate + " exists"
                if !valid_map[row] && !valid_map[col] && !valid_map[area] {
                    valid_map[row] = true
                    valid_map[col] = true
                    valid_map[area] = true
                }else{
                    return false
                }
            }
        }
    }
    return true
} 

/* 
	Example 1:

	Input:
	[
	["5","3",".",".","7",".",".",".","."],
	["6",".",".","1","9","5",".",".","."],
	[".","9","8",".",".",".",".","6","."],
	["8",".",".",".","6",".",".",".","3"],
	["4",".",".","8",".","3",".",".","1"],
	["7",".",".",".","2",".",".",".","6"],
	[".","6",".",".",".",".","2","8","."],
	[".",".",".","4","1","9",".",".","5"],
	[".",".",".",".","8",".",".","7","9"]
	]
	Output: true

	Example 2:

	Input:
	[
	["8","3",".",".","7",".",".",".","."],
	["6",".",".","1","9","5",".",".","."],
	[".","9","8",".",".",".",".","6","."],
	["8",".",".",".","6",".",".",".","3"],
	["4",".",".","8",".","3",".",".","1"],
	["7",".",".",".","2",".",".",".","6"],
	[".","6",".",".",".",".","2","8","."],
	[".",".",".","4","1","9",".",".","5"],
	[".",".",".",".","8",".",".","7","9"]
	]
	Output: false
	Explanation: Same as Example 1, except with the 5 in the top left corner being 
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
*/