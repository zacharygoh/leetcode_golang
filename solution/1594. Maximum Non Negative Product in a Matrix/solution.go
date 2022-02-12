func maxProductPath(grid [][]int) int {
    row, col := len(grid), len(grid[0])
    dp_max := make([][]int, row)
    dp_min := make([][]int, row)
    
    for i := 0; i < row; i++ {
        dp_max[i] = make([]int, col)
        dp_min[i] = make([]int, col)
        for c := 0; c < col; c++ {
            dp_max[i][c] = grid[i][c]
            dp_min[i][c] = grid[i][c]
        }
    }
    
    for r := 1; r < row; r++ {
        dp_max[r][0] = grid[r][0]*dp_max[r-1][0]
        dp_min[r][0] = grid[r][0]*dp_max[r-1][0]
    }
    
    for c := 1; c < col; c++ {
        dp_max[0][c] = grid[0][c]*dp_min[0][c-1]
        dp_min[0][c] = grid[0][c]*dp_min[0][c-1]
    }
    
    fmt.Println(dp_max)
    fmt.Println(dp_min)
    
    for r := 1; r < row; r++ {
        for c := 1; c < col; c++ {
            grid := grid[r][c]
            max1, max2 := dp_max[r-1][c]*grid, dp_max[r][c-1]*grid
            min1, min2 := dp_min[r-1][c]*grid, dp_min[r][c-1]*grid
            dp_max[r][c] = max(max(max1, max2), max(min1, min2))
            dp_min[r][c] = min(min(max1, max2), min(min1, min2))
        }
    }
    if dp_max[row-1][col-1] < 0 {
        return -1
    }
    mod := int(int(math.Pow(10, 9))+int(7))
    return dp_max[row-1][col-1]%mod
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

/* 
Example 1:
Input: grid = [[-1,-2,-3],
               [-2,-3,-3],
               [-3,-3,-2]]
Output: -1
Explanation: It's not possible to get non-negative product in the path from (0, 0) to (2, 2), so return -1.

Example 2:
Input: grid = [[1,-2,1],
               [1,-2,1],
               [3,-4,1]]
Output: 8
Explanation: Maximum non-negative product is in bold (1 * 1 * -2 * -4 * 1 = 8).

Example 3:
Input: grid = [[1, 3],
               [0,-4]]
Output: 0
Explanation: Maximum non-negative product is in bold (1 * 0 * -4 = 0).

Example 4:
Input: grid = [[ 1, 4,4,0],
               [-2, 0,0,1],
               [ 1,-1,1,1]]
Output: 2
Explanation: Maximum non-negative product is in bold (1 * -2 * 1 * -1 * 1 * 1 = 2).
*/