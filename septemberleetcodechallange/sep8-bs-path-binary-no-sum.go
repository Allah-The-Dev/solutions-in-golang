package septemberleetcodechallange

import (
	"fmt"

	"solutions-in-golang/model"
)

func recursiveTraversal(curNo int, listOfPath *[]int, root *model.TreeNode) {
	if root == nil {
		return
	}
	curNo <<= 1
	curNo += root.Val
	if root.Left == nil && root.Right == nil {
		*listOfPath = append(*listOfPath, curNo)
		return
	} else {
		recursiveTraversal(curNo, listOfPath, root.Left)
		recursiveTraversal(curNo, listOfPath, root.Right)
	}
}

func SumRootToLeaf(root *model.TreeNode) int {
	listOfPath := []int{}
	recursiveTraversal(0, &listOfPath, root)
	fmt.Println(listOfPath)
	result := 0
	for _, num := range listOfPath {
		result += num
	}
	return result
}
