/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 { return head }
    rotationCount, nodesInCurrRotation := 0, 0
    node := head
    var prevOfFirstNode, firstNode, lastNode *ListNode
    var prev *ListNode
    listLen := 0
    for node != nil {
        listLen++
        node = node.Next
    }
    totalRotationCount := listLen/k
    node = head
    for node != nil {
        nodesInCurrRotation++
        if nodesInCurrRotation == 1 {
            prev = node
            firstNode = node
            node = node.Next
            continue
        }
        temp := node.Next
        node.Next = prev
        if nodesInCurrRotation == k {
            lastNode = node
        }
        prev = node
        node = temp
        if nodesInCurrRotation == k {
            rotationCount++
            if prevOfFirstNode != nil {
                prevOfFirstNode.Next = lastNode
            }
            firstNode.Next = node
            if rotationCount == 1 {
                head = lastNode
            }
            if rotationCount == totalRotationCount {
                break
            }
            nodesInCurrRotation = 0
            prevOfFirstNode = firstNode
            prev = nil
        }
    }
    return head
}