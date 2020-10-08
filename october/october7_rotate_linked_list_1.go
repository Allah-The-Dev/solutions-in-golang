package october

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	lengthOfLL := 0
	curr := head
	for curr != nil {
		lengthOfLL++
		curr = curr.Next
	}

	rotate := func(head *ListNode) {
		//hold the value of head in a temp var
		cur := head
		currVal := cur.Val
		for cur != nil {
			nextVal := curr.Val
			curr.Val = currVal
			currVal = nextVal

			curr = curr.Next
		}
		head.Val = currVal
	}

	for i := 1; i <= k%lengthOfLL; i++ {
		rotate(head)
	}

	return head
}
