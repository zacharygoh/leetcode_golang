/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    result := new(ListNode)
    out := result
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            result.Next = l1
            l1 = l1.Next
        }else{
            result.Next = l2
            l2 = l2.Next
        }
        result = result.Next
    }
    if l1 != nil {
        result.Next = l1
    }else if l2 != nil {
        result.Next = l2
    }
    return out.Next
}

/* 
	Example 1:
	Input: l1 = [1,2,4], l2 = [1,3,4]
	Output: [1,1,2,3,4,4]

	Example 2:

	Input: l1 = [], l2 = []
	Output: []
	
	Example 3:

	Input: l1 = [], l2 = [0]
	Output: [0]
*/