func isMonotonic(A []int) bool {
    is_increasing := true
    if A[0] < A[len(A)-1] {
        is_increasing = true
    }else if A[0] > A[len(A)-1] {
        is_increasing = false
    }
    for i := 1; i < len(A); i++ {
        if A[i - 1] <= A[i] && is_increasing {
            continue
        }else if A[i - 1] >= A[i] && !is_increasing {
            continue
        }else{
            return false
        }
    }
    return true
}

/*
Example 1:

Input: [1,2,2,3]
Output: true

Example 2:

Input: [6,5,4,4]
Output: true

Example 3:

Input: [1,3,2]
Output: false

Example 4:

Input: [1,2,4,5]
Output: true

Example 5:

Input: [1,1,1]
Output: true 
*/