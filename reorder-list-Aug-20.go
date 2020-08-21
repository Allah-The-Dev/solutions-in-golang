package main

import "fmt"

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	prev = nil
	curr := head
	next := head.Next
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var next *ListNode
	for l2 != nil {
		next = l1.Next
		l1.Next = l2
		l1 = l2
		l2 = next
	}
}

func reorderList(head *ListNode) {
	//boundary condition
	//if list is null or have one item
	if head == nil || head.Next == nil {
		return
	}
	//break the list in 2 parts
	//using slow fast approach
	prev := head
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil //this will break list in 2 parts //one before slow one from slow
	//reverse the second list
	l2 := reverse(slow)
	//merge both the list
	mergeList(head, l2)
}

func printList(head *ListNode) {
	fmt.Printf("%d -> ", head.Val)
	head = head.Next
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Printf("%d -> ", list.Val)
	printList(list.Next)
	reorderList(list)
}
