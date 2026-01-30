/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reverseList(head *ListNode) *ListNode {
	if head == nil{return nil}
    prev:=&ListNode{}
	for head != nil{
		temp:=head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}