package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge_sort(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var prev *ListNode
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    prev.Next = nil // указатель перед средним эелментом ставим в ноль
    l := merge_sort(head)
    r := merge_sort(slow)

    dummy := &ListNode{}
    curr := dummy
    
    for l != nil && r != nil {
        if l.Val < r.Val {
            curr.Next = l
            l = l.Next
        } else {
            curr.Next = r
            r = r.Next
        }
        curr = curr.Next
    }
    if l != nil {
        curr.Next = l
    } else {
        curr.Next = r
    }

    return dummy.Next
}
func sortList(head *ListNode) *ListNode {
	return merge_sort(head)
    
}