package main


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head:=&ListNode{Val:0, Next:nil}
	cur:= head
	left:= list1
	right:= list2
	for left!=nil && right !=nil{
		if left.Val <= right.Val{
			cur.Next = &ListNode{Val:left.Val}
			cur = cur.Next
			left = left.Next
		}else if right.Val<=left.Val{
			cur.Next = &ListNode{Val:right.Val}
			cur = cur.Next
			right = right.Next

		}
	}
	if left == nil && right == nil{
		return head.Next
	}else if left==nil && right!=nil{
		for right !=nil{
			cur.Next = &ListNode{Val:right.Val}
			cur = cur.Next
			right = right.Next
		}
	}else if left!=nil && right == nil{
		for left!=nil{
			cur.Next = &ListNode{Val:left.Val}
			cur = cur.Next
			cur = cur.Next
			left = left.Next
		}
	}
	return head.Next
}