/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    
    sum, carryOver := 0, false
    
    sum, carryOver = getSumAndCarryOver(l1.Val, l2.Val, carryOver)
    l1, l2 = l1.Next, l2.Next
    
    res = &ListNode{sum, nil}
    resN := res
    
    for l1 != nil || l2 != nil {
        prevCO := carryOver
        if l1 != nil && l2 != nil {
            sum, carryOver = getSumAndCarryOver(l1.Val, l2.Val, prevCO)
            l1, l2 = l1.Next, l2.Next
        } else if l1 != nil {
            sum, carryOver = getSumAndCarryOver(l1.Val, 0, prevCO)
            l1 = l1.Next
        } else if l2 != nil {
            sum, carryOver = getSumAndCarryOver(0, l2.Val, prevCO)
            l2 = l2.Next
        }
        resN.Next = &ListNode{sum, nil}
        resN = resN.Next
    }
    if carryOver {
        resN.Next = &ListNode{1, nil}
    }
    return res
}

func getSumAndCarryOver(a, b int, co bool) (int, bool) {
    currSum := a+b
    if co {
        currSum = a+b+1
    }
    if currSum >= 10 {
        return currSum%10, true
    }
    return currSum, false
}