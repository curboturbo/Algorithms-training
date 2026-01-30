package main
import "fmt"


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right || head.Next == nil {
        return head
    }
    root := &ListNode{Next: head}
    dummy := root
    pre := head

    for i := 1; i < left; i++ {
        dummy = pre
        pre = pre.Next
    }

    pointer := pre
    save_pointer := pre
    var before *ListNode

    for i := left; i < right; i++ {
        tmp := pointer.Next
        pointer.Next = before
        before = pointer
        pointer = tmp
    }
    up := pointer.Next
    pointer.Next = before
    
    save_pointer.Next = up
    dummy.Next = pointer

    return root.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
