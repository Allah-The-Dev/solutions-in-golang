package october

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	} else if val > root.Val {
		root.Right = InsertIntoBST(root.Right, val)
	} else {
		root.Left = InsertIntoBST(root.Left, val)
	}

	return root
}
