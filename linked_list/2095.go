package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    smart:= head
	slow:=head
	prev:=&ListNode{}
	if head == nil{return nil}
	if head.Next == nil{return nil}
	for smart != nil{
		smart = smart.Next
		if smart == nil{break}
		smart = smart.Next
		prev=slow
		slow = slow.Next
	}
	// slow указывает на удаляемый
	
	temp:=slow.Next
	slow = nil
	prev.Next = temp
	return head
}