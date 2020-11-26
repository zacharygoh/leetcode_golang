/** 
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

 func findSolution(customFunction func(int, int) int, z int) (result [][]int) {
    for i := 1; i <= z; i++ {
        for j := 1; j <= z; j++ {
            if customFunction(i, j) == z {
                result = append(result, []int{i, j})
            }
        }
    }
    return
}

/* 
	Example 1:

	Input: function_id = 1, z = 5
	Output: [[1,4],[2,3],[3,2],[4,1]]
	Explanation: function_id = 1 means that f(x, y) = x + y

	Example 2:

	Input: function_id = 2, z = 5
	Output: [[1,5],[5,1]]
	Explanation: function_id = 2 means that f(x, y) = x * y
*/