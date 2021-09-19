/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    return isValueAtNodeInLimit(root, math.MinInt64, math.MaxInt64)
}


// approach - 1
func isValueAtNodeInLimit(root *TreeNode, minLimit, maxLimit int) bool {
    if root == nil { return true } 
    
    if root.Val > minLimit && root.Val < maxLimit {
        return isValueAtNodeInLimit(root.Left, minLimit, root.Val) && isValueAtNodeInLimit(root.Right, root.Val, maxLimit)
    } else {
        return false
    }
}