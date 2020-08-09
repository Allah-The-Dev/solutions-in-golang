package main

import "fmt"

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumUtil(root *TreeNode, dataSumMap map[int]int, cSum int, sum int, result *int) {
	if root == nil {
		return
	}
	temp := cSum + root.Val
	if val, ok := dataSumMap[temp-sum]; ok {
		*result += val
	}

	if val, ok := dataSumMap[temp]; ok {
		dataSumMap[temp] = val + 1
	} else {
		dataSumMap[temp] = 1
	}
	//fmt.Printf("before node %v\n", dataSumMap)
	if root.Left != nil {
		pathSumUtil(root.Left, dataSumMap, cSum+root.Val, sum, result)
	}
	if root.Right != nil {
		pathSumUtil(root.Right, dataSumMap, cSum+root.Val, sum, result)
	}
	dataSumMap[temp] = dataSumMap[temp] - 1
	//fmt.Printf("after node %v\n", dataSumMap)
}

func pathSum(root *TreeNode, sum int) int {

	dataSumMap := map[int]int{0: 1}
	result := 0

	pathSumUtil(root, dataSumMap, 0, sum, &result)

	return result

}

func main() {
	bt := &TreeNode{}
	bt.Val = 10

	bt.Left = &TreeNode{}
	bt.Left.Val = 5

	bt.Right = &TreeNode{}
	bt.Right.Val = -3

	bt.Right.Right = &TreeNode{}
	bt.Right.Right.Val = 11

	bt.Left.Left = &TreeNode{3, &TreeNode{3, nil, nil}, &TreeNode{-2, nil, nil}}

	bt.Left.Right = &TreeNode{2, nil, &TreeNode{1, nil, nil}}

	fmt.Printf("%v", pathSum(bt, 8))
}
