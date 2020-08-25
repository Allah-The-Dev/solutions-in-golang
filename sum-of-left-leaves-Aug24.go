package main

import (
	"container/list"
	"fmt"
)

//TreeNode ...
// type TreeNodeT struct {
//     Val int
// 	Left *TreeNodeT
// 	Right *TreeNodeT
// }

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := list.New()
	leftLeafSum := 0
	current := root
	for {
		for current != nil {
			stack.PushFront(current)
			leftOfCurrent := current.Left
			if leftOfCurrent != nil && leftOfCurrent.Left == nil && leftOfCurrent.Right == nil {
				leftLeafSum += leftOfCurrent.Val
				current = nil
			} else {
				current = current.Left
			}
		}
		current = stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		current = current.Right
		if current == nil && stack.Len() == 0 {
			break
		}
	}
	return leftLeafSum
}

func main() {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(sumOfLeftLeaves(tree))
}
