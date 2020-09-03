func validMountainArray(A []int) bool {
    is_increasing := true
    largest_num := 0
    if len(A) == 0 || len(A) == 1{
        return false
    }
    for i := 1; i < len(A); i++ {
        if A[i - 1] < A[i] && is_increasing {
            largest_num = A[i]
            if i + 1 < len(A) && A[i + 1] < largest_num {
                is_increasing = false
            }
        }else if A[i - 1] > A[i] && !is_increasing {
            continue
        }else{
            return false
        }
    }
    if A[len(A)-2] < A[len(A)-1] {
        return false
    }
    return true
}

/* 
Example 1:

Input: [2,1]
Output: false

Example 2:

Input: [3,5,5]
Output: false

Example 3:

Input: [0,3,2,1]
Output: true
*/