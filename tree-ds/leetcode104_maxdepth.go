/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    return dfs(root, 1)
}

func dfs(root *TreeNode, currDepth int) int {
    //if root == nil { return 0 }
    
    leftDepth, rightDepth := currDepth, currDepth
    
    if root.Left != nil { 
        leftDepth = dfs(root.Left, leftDepth+1)
    }
    
     if root.Right != nil { 
        rightDepth = dfs(root.Right, rightDepth+1)
    }
    
    return max(leftDepth, rightDepth)
}

func max(a, b int) int {
    if a > b { return a }
    return b
}