/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    value := 0
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        value = root.Left.Val
    }
    
    return value + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

/* 
	Example:

	 3
	/ \
   9  20
     /  \
	15   7

	There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/