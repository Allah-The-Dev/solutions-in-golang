/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// each pair will have value of list node, and pointer to that node
// using which list can be traversed further

// idea is that take first node of all the list and put it in 
// min priority queue
// so that min always stays in top

type pair struct {
    Val int
    node *ListNode
}

type pairs []*pair

func (ps pairs) Len() int { return len(ps) }

func (ps pairs) Less(i, j int) bool { return ps[i].Val < ps[j].Val }

func (ps pairs) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

func (ps *pairs) Push(elem interface{}) { *ps = append(*ps, elem.(*pair)) }

func (ps *pairs) Pop() interface{} {
    var elem *pair
    psLen := len(*ps)
    elem, *ps, (*ps)[psLen-1] = (*ps)[psLen-1], (*ps)[:psLen-1], nil
    return elem
}
 
func mergeKLists(lists []*ListNode) *ListNode {
    // using priority queue
    var res *ListNode
    
    listLen := len(lists)
    if listLen == 0 { return res }
    
    // priority queue to main min at top
    pq := pairs{}
    
    // take first element from each list and 
    for _, list := range lists {
        if list != nil {
            pq = append(pq, &pair{list.Val, list})
        }
    }
    
    if len(pq) == 0 { return res }
    
    heap.Init(&pq)
    
    minHeapTop := heap.Pop(&pq).(*pair)
    res = &ListNode{minHeapTop.Val, nil}
    if minHeapTop.node.Next != nil {
        heap.Push(&pq, &pair{minHeapTop.node.Next.Val, minHeapTop.node.Next})
    }
    
    resN := res
    
    for len(pq) > 0 {
        minHeapTop = heap.Pop(&pq).(*pair)
        resN.Next = &ListNode{minHeapTop.Val, nil}
        resN = resN.Next
        if minHeapTop.node.Next != nil {
            heap.Push(&pq, &pair{minHeapTop.node.Next.Val, minHeapTop.node.Next})
        }
    }
    
    return res
}