/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func binaryTreePaths(root *TreeNode) (result []string) {
    if root == nil {
        return nil
    }
    helper(root, strconv.Itoa(root.Val), &result)
    return
}

func helper(root *TreeNode, str string, result *[]string) {
    if root.Left == nil && root.Right == nil {
        *result = append(*result, str)
        return
    }
    
    if root.Left != nil {
        helper(root.Left, str + "->" + strconv.Itoa(root.Left.Val), result)
    }
    if root.Right != nil {
        helper(root.Right, str + "->" + strconv.Itoa(root.Right.Val), result)
    }
}

/* 
	Example:

	Input:

	1
	/   \
	2     3
	\
	5

	Output: ["1->2->5", "1->3"]

	Explanation: All root-to-leaf paths are: 1->2->5, 1->3
*/