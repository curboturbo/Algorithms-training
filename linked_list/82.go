package main

func deleteDuplicates(head *ListNode) *ListNode {
    ptr := &ListNode{Next: head}
	pre:=ptr
	for head != nil && head.Next != nil{
		if head.Next != nil && head.Val==head.Next.Val{
			for head.Next!=nil && head.Val==head.Next.Val{
				head = head.Next
			}
			head = head.Next
			pre.Next=head
		}else{
			pre = head
			head = head.Next
		}
	}
	return ptr.Next
}
