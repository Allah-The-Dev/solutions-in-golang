/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    listOfVal := []int{}
    inorder(root, &listOfVal)
    isSorted := sort.IntsAreSorted(listOfVal)
    if !isSorted {
        return false
    }
    //fmt.Println(listOfVal)
    areDuplicates := checkDuplicates(listOfVal)
    return !areDuplicates
}

func inorder(root *TreeNode, nums *[]int) {
    if root == nil { return }
    
    inorder(root.Left, nums)
    *nums = append(*nums, root.Val)
    inorder(root.Right, nums)
}

func checkDuplicates(nums []int) bool {
    res := false
    for i, num := range nums {
        if i == 0 { continue }
        if num == nums[i-1] {
            res = true
            break
        }
    }
    return res
}