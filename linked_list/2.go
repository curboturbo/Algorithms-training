package main
import "fmt"



func print(list *ListNode){
	for list != nil{
		fmt.Printf("%d ",list.Val)
		list = list.Next
	}
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{Val: 0}
    curr := head
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        val1, val2 := 0, 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }

        sum := val1 + val2 + carry
        carry = sum / 10
        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
    }
    return head.Next
}

//func main() {

//	l1 := &ListNode{Val: 2}
//	l1.Next = &ListNode{Val: 4}
//	l1.Next.Next = &ListNode{Val: 3}

//	l2 := &ListNode{Val: 5}
//	l2.Next = &ListNode{Val: 6}
//	l2.Next.Next = &ListNode{Val: 4}
//
//	result := addTwoNumbers(l1, l2)
//
//	fmt.Print("Результат: ")
//	for result != nil {
//		fmt.Printf("%d ", result.Val)
//		result = result.Next
//	}
//}