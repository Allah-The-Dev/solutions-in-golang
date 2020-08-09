package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//TreeNode ...
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

//TreeNodeWithHD ...
type TreeNodeWithHD struct {
	Node *TreeNode
	HD   int
}

func getHDLevelOrderedMap(root *TreeNode, hdMap map[int][]int, hd int, hdMin, hdMax *int) {

	if root == nil {
		return
	}

	queue := []*TreeNodeWithHD{{root, 0}}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		//update map
		if nodeDataArr, ok := hdMap[node.HD]; ok {
			nodeDataArr = append(nodeDataArr, node.Node.Val)
			hdMap[node.HD] = nodeDataArr
		} else {
			nodeDataArr = []int{node.Node.Val}
			hdMap[node.HD] = nodeDataArr
		}
		if *hdMin > node.HD {
			*hdMin = node.HD
		}
		if *hdMax < node.HD {
			*hdMax = node.HD
		}
		if node.Node.Left != nil {
			queue = append(queue, &TreeNodeWithHD{node.Node.Left, node.HD - 1})
		}
		if node.Node.Right != nil {
			queue = append(queue, &TreeNodeWithHD{node.Node.Right, node.HD + 1})
		}
	}

}

func verticalTraversal(root *TreeNode) [][]int {

	hdMap := make(map[int][]int)
	hdMin, hdMax := 0, 0
	getHDLevelOrderedMap(root, hdMap, 0, &hdMin, &hdMax)

	result := make([][]int, len(hdMap))

	for i, j := hdMin, 0; i <= hdMax; i++ {
		result[j] = hdMap[i]
		j++
	}

	return result
}

// func main() {
// 	bt := &TreeNode{}
// 	bt.Val = 0
// 	bt.Left = &TreeNode{}
// 	bt.Left.Val = 8
// 	bt.Right = &TreeNode{}
// 	bt.Right.Val = 1

// 	bt.Right.Left = &TreeNode{}
// 	bt.Right.Left.Val = 3
// 	bt.Right.Right = &TreeNode{}
// 	bt.Right.Right.Val = 2

// 	bt.Right.Left.Right = &TreeNode{4, nil, &TreeNode{7, nil, nil}}

// 	bt.Right.Right.Left = &TreeNode{5, &TreeNode{6, nil, nil}, nil}
// 	fmt.Printf("%v", verticalTraversal(bt))
// }
