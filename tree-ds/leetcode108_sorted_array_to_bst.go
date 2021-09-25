/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    numsLen := len(nums)
    mid := len(nums)/2
    if numsLen == 0 {
        return nil
    }
    root := &TreeNode{nums[mid], nil, nil}
    if numsLen == 1 {
        return root
    }
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    return root
}