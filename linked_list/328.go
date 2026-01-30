/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil || head.Next.Next == nil{return head}
	even:=1
	even_head:=&ListNode{}
	pointer:= even_head
	double_pointer_to_head := head
	prev:= &ListNode{}
	for head!= nil{
		if even % 2 !=0{
			prev = head
			head = head.Next
			even ++
		}else if even % 2 ==0{
			temp:= head.Next
			prev.Next = temp
			head.Next = nil
			even_head.Next = head
			even_head = even_head.Next
			even++
			head = temp
		}
	}
	prev.Next = pointer.Next
	return double_pointer_to_head
}