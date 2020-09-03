func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    if rec2[0] >= rec1[2] || rec1[1] >= rec2[3] {
        return false
    }
    if rec2[2] <= rec1[0] || rec1[3] <= rec2[1] {
        return false
    }
    return true
}

/*

Example 1:

Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
Output: true

Example 2:

Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
Output: false
*/