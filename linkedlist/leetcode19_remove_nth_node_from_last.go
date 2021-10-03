/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// idea is that first move fast pointer to the right
// to create a distance of "n"
// between fast and slow pointer
// so that when fast pointer reaches to the last of linked list 
// slow pointer will be listLen-n node behind
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{-1, head}
    fast, slow := dummy, dummy
    for fast.Next != nil {
        if n <= 0 {
            slow = slow.Next
        }
        fast = fast.Next
        n -= 1
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}