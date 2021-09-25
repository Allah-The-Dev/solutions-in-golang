/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // recursive solution
 func buildTree(preorder []int, inorder []int) *TreeNode {
    
    preorderLen := len(preorder)
    
    if preorderLen == 0 { 
        return nil 
    } 
    
    root := &TreeNode{preorder[0], nil, nil}
    
    if preorderLen == 1 {
        return root
    }
    
    // find index or root element in inorder
    idxOfRoot := -1
    for i, num := range inorder {
        if num == preorder[0] {
            idxOfRoot = i 
            break
        }
    }
    rightPreOrderLen := preorderLen-idxOfRoot-1
    
    root.Left = buildTree(preorder[1:preorderLen-rightPreOrderLen], inorder[:idxOfRoot])
    root.Right = buildTree(preorder[preorderLen-rightPreOrderLen:], inorder[idxOfRoot+1:])
    
    return root
}