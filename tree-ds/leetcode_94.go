// inorder traversal   94
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    inorder(root, &res)
    return res
}

func inorder(root *TreeNode, res *[]int) {
    if root == nil { return }
    inorder(root.Left, res)
    *res = append(*res, root.Val)
    inorder(root.Right, res)
}