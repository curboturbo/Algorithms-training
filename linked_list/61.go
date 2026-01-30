package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    cnt := 0
    ptr := head
    var last *ListNode
    for ptr != nil {
        cnt++
        last = ptr
        ptr = ptr.Next
    }

    rotate := k % cnt
    if rotate == 0 {
        return head
    }
    last.Next = head
    
    ptr = head
    stop_point := cnt - rotate
    current_pos := 1
    for current_pos < stop_point {
        current_pos++
        ptr = ptr.Next
    }

    z := ptr.Next
    ptr.Next = nil
    return z
}


func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}


func Print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	nums := []int{0,1,2}
	head := createList(nums)

	fmt.Println("Linked list before function")
	Print(head)

	// 2. Удаляем дубликаты
	result := rotateRight(head, 4)

	fmt.Println("Linked list after function")
	Print(result)
}
