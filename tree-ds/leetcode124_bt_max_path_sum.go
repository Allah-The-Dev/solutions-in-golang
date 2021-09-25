/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// take left sum , take right sum at any particular node
// set max path sum = max(currMax, left+right+node)
// and this node will return max(left, right)+node
func maxPathSum(root *TreeNode) int {
    maxPS := math.MinInt8
    
    var maxPath func(*TreeNode) int
    maxPath = func(node *TreeNode) int {
        if node == nil { return 0 }
        
        leftSum := max(0, maxPath(node.Left))
        rightSum := max(0, maxPath(node.Right))
        
        maxPS = max(maxPS, leftSum+rightSum+node.Val)
        
        return max(leftSum, rightSum)+node.Val
    }
    
    maxPath(root)
    
    return maxPS
}

func max(a, b int) int {
    if a > b { return a }
    return b
}