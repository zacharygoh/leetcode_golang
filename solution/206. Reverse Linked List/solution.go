/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    head = prev
    return head
}

/* 
	Example 1:

	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

	Example 2:

	Input: head = [1,2]
	Output: [2,1]
	
	Example 3:

	Input: head = []
	Output: []
*/