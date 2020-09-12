package septemberleetcodechallange

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, sortedList *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, sortedList)
	*sortedList = append(*sortedList, root.Val)
	inorder(root.Right, sortedList)
}
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	sortedList := []int{}
	inorder(root1, &sortedList)
	inorder(root2, &sortedList)
	sort.Ints(sortedList)
	return sortedList
}
