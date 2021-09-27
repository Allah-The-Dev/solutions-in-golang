/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    var res *ListNode
    if len(lists) == 0 { return res}
    
    for len(lists) > 1 {
        tempList := []*ListNode{}
        for i := 0; i < len(lists); i += 2 {
            if i+1 < len(lists) {
                mergedList := merge2Lists(lists[i], lists[i+1])
                tempList = append(tempList, mergedList)   
            } else {
                if lists[i] != nil {
                    tempList = append(tempList, lists[i])
                }
            }
        }
        lists = tempList
    }
    return lists[0]
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
    var res *ListNode
    
    if list1 == nil { return list2 }
    if list2 == nil { return list1 }
    
    if list1.Val < list2.Val {
        res = &ListNode{list1.Val, nil}
        list1 = list1.Next
    } else {
        res = &ListNode{list2.Val, nil}
        list2 = list2.Next
    }
    
    resN := res
    
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            resN.Next = &ListNode{list1.Val, nil}
            resN = resN.Next
            list1 = list1.Next
        } else {
            resN.Next = &ListNode{list2.Val, nil}
            resN = resN.Next
            list2 = list2.Next
        }
    }
    
    if list1 != nil {
        resN.Next = list1
    } else if list2 != nil {
        resN.Next = list2
    }
    
    return res
}