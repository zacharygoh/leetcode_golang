/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var temp *ListNode
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        temp = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    temp.Next = nil
    
    left_side, right_side := sortList(head), sortList(slow)
    
    return merge(left_side, right_side)
    return head
}

func merge(left, right *ListNode) *ListNode {
    var head *ListNode
    var curr *ListNode
    
    if left.Val < right.Val {
        head = left
        left = left.Next
    }else{
        head = right
        right = right.Next
    }
    curr = head
    
    for left != nil && right != nil {
        if left.Val < right.Val {
            curr.Next = left
            left = left.Next
        }else{
            curr.Next = right
            right = right.Next
        }
        curr = curr.Next
    }
    
    if left != nil {
        curr.Next = left
    }else if right != nil {
        curr.Next = right
    }
    
    return head
}

/* 
	Example 1:

	Input: head = [4,2,1,3]
	Output: [1,2,3,4]

	Example 2:

	Input: head = [-1,5,3,4,0]
	Output: [-1,0,3,4,5]

	Example 3:

	Input: head = []
	Output: []
*/