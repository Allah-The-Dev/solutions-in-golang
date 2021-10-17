/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// approach - 1 : using hashmap of node and copy node
 func copyRandomList(head *Node) *Node {
    var res *Node
    if head == nil {
        return res
    }
    nodeMap := make(map[*Node]*Node)
    node := head 
    for node != nil {
        nodeMap[node] = &Node{node.Val, nil, nil}
        node = node.Next
    }
    
    res = nodeMap[head]
    resHead := res
    
    node = head
    for node != nil {
        resHead = nodeMap[node]
        resHead.Next = nodeMap[node.Next]
        resHead.Random = nodeMap[node.Random]
        node = node.Next
        resHead = resHead.Next
    }
    
    return res
}

// approach - 2 : use temporary copy node
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
    var res *Node
    if head == nil {
        return res
    }
    node := head
    // insert copy of nodes to next of original node
    // and reassign pointers
    for node != nil {
        nodeNext := node.Next
        copyOfNode := &Node{node.Val, nodeNext, nil}
        node.Next = copyOfNode
        node = nodeNext
    }
    // assign random pointers 
    // using mechanism of node.next.random = node.random.next
    node = head
    for node != nil && node.Next != nil {
        if node.Random == nil {
            node.Next.Random = nil 
        } else {
            node.Next.Random = node.Random.Next
        }
        node = node.Next.Next
    }
    // now remove nodes of original nodes 
    node = head
    res = head.Next
    resHead := res
    for resHead != nil && resHead.Next != nil {
        node.Next = resHead.Next
        node = node.Next
        
        resHead.Next = node.Next
        resHead = resHead.Next
    }
    node.Next = resHead.Next
    return res
}

/*

node = head 
fmt.Println()
for node != nil {
    fmt.Printf("%d -> ", node.Val)
    node = node.Next
}

 node = head 
for node != nil {
    if node.Random != nil {
        fmt.Printf("%d -> ", node.Random.Val)   
    } else {
        fmt.Print("nil -> ")   
    }
    node = node.Next
}

*/