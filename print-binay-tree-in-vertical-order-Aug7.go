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

//TreeNodeWithHDAndVD ...
type TreeNodeWithHDAndVD struct {
	Node *TreeNode
	HD   int
	VD   int
}

//TreeDataWithHDAndVD ...
type TreeDataWithHDAndVD struct {
	val, HD, VD int
}

func getHDLevelOrderedMap(root *TreeNode, hdMap map[int][]TreeDataWithHDAndVD, hd int, hdMin, hdMax *int) {

	if root == nil {
		return
	}

	queue := []*TreeNodeWithHDAndVD{{root, 0, 0}}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		//update map
		if nodeDataArr, ok := hdMap[node.HD]; ok {
			hdMap[node.HD] = addNodeAsPerVDOrderAndSorted(nodeDataArr, node)
		} else {
			nodeDataArr = []TreeDataWithHDAndVD{{node.Node.Val, node.HD, node.VD}}
			hdMap[node.HD] = nodeDataArr
		}
		if *hdMin > node.HD {
			*hdMin = node.HD
		}
		if *hdMax < node.HD {
			*hdMax = node.HD
		}
		if node.Node.Left != nil {
			queue = append(queue, &TreeNodeWithHDAndVD{node.Node.Left, node.HD - 1, node.VD + 1})
		}
		if node.Node.Right != nil {
			queue = append(queue, &TreeNodeWithHDAndVD{node.Node.Right, node.HD + 1, node.VD + 1})
		}
	}

}

func addNodeAsPerVDOrderAndSorted(nodeDataArr []TreeDataWithHDAndVD, node *TreeNodeWithHDAndVD) []TreeDataWithHDAndVD {
	appended := false
	for i, savedNodeData := range nodeDataArr {
		if savedNodeData.VD == node.VD {
			if node.Node.Val < savedNodeData.val {
				nodeDataArr = append(nodeDataArr, TreeDataWithHDAndVD{})
				copy(nodeDataArr[i+1:], nodeDataArr[i:])
				nodeDataArr[i] = TreeDataWithHDAndVD{node.Node.Val, node.HD, node.VD}
				appended = true
				break
			} else if node.Node.Val > savedNodeData.val && i+1 <= len(nodeDataArr)-1 && nodeDataArr[i+1].VD != node.VD {
				nodeDataArr = append(nodeDataArr, TreeDataWithHDAndVD{})
				copy(nodeDataArr[i+2:], nodeDataArr[i+1:])
				nodeDataArr[i+1] = TreeDataWithHDAndVD{node.Node.Val, node.HD, node.VD}
				appended = true
				break
			}
		}
	}
	if !appended {
		nodeDataArr = append(nodeDataArr, TreeDataWithHDAndVD{node.Node.Val, node.HD, node.VD})
	}
	return nodeDataArr
}

func verticalTraversal(root *TreeNode) [][]int {

	hdMap := make(map[int][]TreeDataWithHDAndVD)
	hdMin, hdMax := 0, 0
	getHDLevelOrderedMap(root, hdMap, 0, &hdMin, &hdMax)

	result := make([][]int, len(hdMap))

	for i, j := hdMin, 0; i <= hdMax; i++ {
		resultOfHD := make([]int, len(hdMap[i]))
		for k, treeDataPoint := range hdMap[i] {
			resultOfHD[k] = treeDataPoint.val
		}
		result[j] = resultOfHD
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
