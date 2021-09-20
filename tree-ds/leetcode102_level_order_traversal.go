/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    
    if root == nil { return res }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        queueLen := len(queue)
        tempRes := make([]int, queueLen)
        for i := 0; i < queueLen; i++ {
            queueTop := queue[0]
            queue = queue[1:]
            tempRes[i] = queueTop.Val
            if queueTop.Left != nil {
                queue = append(queue, queueTop.Left)
            }
            if queueTop.Right != nil {
                queue = append(queue, queueTop.Right)
            }
        }
        res = append(res, tempRes)
    }
    return res
}