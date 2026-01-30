package main


//head = [1,4,3,2,5,2], x = 3
//Output: [1,2,2,4,3,5]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
   pointer_head:=&ListNode{}
   pointer_tail:=&ListNode{}
   less := pointer_head
   great:= pointer_tail
   for head!=nil{
	if head.Val < x{
		less.Next = head
		less = less.Next
	}else{
		great.Next = head
		great = great.Next
	}
    head = head.Next
   } 
   less.Next = pointer_tail.Next
   return pointer_head.Next
}

