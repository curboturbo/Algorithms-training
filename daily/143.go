package main



type ListNode struct {	
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil{return}
    slow:=head
    fast:=head
    for fast != nil && fast.Next!=nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    new_head:=slow.Next
    slow.Next = nil
    var prev *ListNode

    for new_head != nil{
		next:=new_head.Next
		new_head.Next = prev
		prev = new_head
		new_head = next
    }
    first := head
    for prev != nil {
    n1 := first.Next
    n2 := prev.Next
    first.Next = prev
    prev.Next = n1
    first = n1
    prev = n2
    }
}