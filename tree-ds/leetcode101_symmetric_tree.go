/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}

func isMirror(rootA, rootB *TreeNode) bool {
    if rootA == nil && rootB == nil { return true }
    if rootA == nil || rootB == nil { return false }
    
    if rootA.Val != rootB.Val { return false }
    
    return isMirror(rootA.Left, rootB.Right) && isMirror(rootA.Right, rootB.Left)
}