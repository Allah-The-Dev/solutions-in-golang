package main

import "fmt"

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

func getHDLevelOrderedMap(root *TreeNode, hdMap map[int][]int, hd int, hdMin, hdMax *int) {

	if root == nil {
		return
	}

	if hdLevelArr, ok := hdMap[hd]; ok {
		for i, num := range hdLevelArr {
			if len(hdLevelArr) == 1 {
				if root.Val < num {
					fmt.Println(hdLevelArr)
					hdLevelArr = append(hdLevelArr, 0)
					copy(hdLevelArr[1:], hdLevelArr[:])
					hdLevelArr[0] = root.Val
				} else {
					hdLevelArr = append(hdLevelArr, 0)
					copy(hdLevelArr[0:], hdLevelArr[:])
					hdLevelArr[1] = root.Val
				}
			} else {
				if root.Val < num {
					hdLevelArr = append(hdLevelArr, 0)
					copy(hdLevelArr[i+1:], hdLevelArr[i:])
					hdLevelArr[i] = root.Val
				}
				if root.Val > num && i == len(hdLevelArr)-1 {
					hdLevelArr = append(hdLevelArr, root.Val)
				}
			}
		}
		hdMap[hd] = hdLevelArr
	} else {
		newHDLevelArr := []int{root.Val}
		hdMap[hd] = newHDLevelArr
	}

	if *hdMin > hd {
		*hdMin = hd
	}
	if *hdMax < hd {
		*hdMax = hd
	}

	getHDLevelOrderedMap(root.Left, hdMap, hd-1, hdMin, hdMax)
	getHDLevelOrderedMap(root.Right, hdMap, hd+1, hdMin, hdMax)

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

func main() {
	bt := &TreeNode{}
	bt.Val = 0
	bt.Left = &TreeNode{}
	bt.Left.Val = 8
	bt.Right = &TreeNode{}
	bt.Right.Val = 1

	bt.Right.Left = &TreeNode{}
	bt.Right.Left.Val = 3
	bt.Right.Right = &TreeNode{}
	bt.Right.Right.Val = 2

	bt.Right.Left.Right = &TreeNode{4, nil, &TreeNode{7, nil, nil}}

	bt.Right.Right.Left = &TreeNode{5, &TreeNode{6, nil, nil}, nil}
	fmt.Printf("%v", verticalTraversal(bt))
}
