/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // ---------------- reverse post order traversal --------------------
 // right -> left -> root
 // remembering prev node so that it can be assigned to right of prev call in stack
 func flatten(root *TreeNode)  {
    var prev *TreeNode
    prev = nil
    
    var reversePostOrder func(*TreeNode)
    reversePostOrder = func(node *TreeNode) {
        if node == nil {
            return
        }
        reversePostOrder(node.Right)
        reversePostOrder(node.Left)

        node.Right = prev
        node.Left = nil
        prev = node
    }
    
    reversePostOrder(root)
}

// ----------------------------- using stack approach ----------------------
// put right then left to stack
//        1
//       / \
//      2   3
// stack [1]
// stack [3, 2]
// stack [3] 2 will be popped out and 2's right = stack[top], 2's left = nil
// stack [] 3 will be popped out
func flatten(root *TreeNode)  {
    stack := []*TreeNode{root}
    var stackTop *TreeNode
    for len(stack) > 0 {
        stackLen := len(stack)
        stackTop, stack = stack[stackLen-1], stack[:stackLen-1]
        if stackTop.Right != nil {
            stack = append(stack, stackTop.Right)
        }
        if stackTop.Left != nil {
            stack = append(stack, stackTop.Left)
        }
        if len(stack) > 0 { 
            stackTop.Right = stack[len(stack)-1]
        }
        stackTop.Left = nil
    }
}

