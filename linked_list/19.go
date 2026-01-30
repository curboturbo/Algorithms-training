package main



//func printList(head *ListNode) {
//	for head != nil {
//		fmt.Printf("%d -> ", head.Val)
//		head = head.Next
//	}
//	fmt.Println("nil")
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    cnt := 0
    ptr := head
    for ptr != nil {
        cnt++
        ptr = ptr.Next
    }
    before_node := dummy
    for i := 0; i < cnt-n; i++ {
        before_node = before_node.Next
    }

    if before_node.Next != nil {
        before_node.Next = before_node.Next.Next
    }
    return dummy.Next
}