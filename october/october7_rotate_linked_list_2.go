package october

func rotateRight2(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	//create a cycle first
	lenOfLL := 1
	last := head
	for last.Next != nil {
		lenOfLL++
		last = last.Next
	}
	last.Next = head

	//if len = 5 and rotatio = 8
	//then actual rotation = 8 % 5 = 3
	//now iterate 2 times and return the node pointer of next i.e. 3rd
	cur := head
	for i := 1; i < lenOfLL-(k%lenOfLL); i++ {
		cur = cur.Next
	}
	res := cur.Next
	cur.Next = nil
	return res
}
