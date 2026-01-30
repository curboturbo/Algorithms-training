/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main
import "math"

func find_max(x int, y int) int{
	if x>=y{return x}else{return y}
}

func pairSum(head *ListNode) int {
   if head == nil{return 0}
   if head.Next == nil {return head.Val}
   if head.Next.Next == nil{return head.Val +head.Next.Val}
   answer:=math.MinInt32
   start := head
   slow:=head
   fast:=head
   prev:=&ListNode{}
   for fast != nil{
	fast = fast.Next
	prev = slow
	slow = slow.Next
	if fast == nil{break}
	fast = fast.Next
   }
   for slow != nil{
	temp:= slow.Next
	slow.Next = prev
	prev = slow
	slow = temp
   }
   end:= prev 
   for start.Next!=end{
	answer = find_max(answer,end.Val+start.Val)
	end = end.Next
	start = start.Next
   }
   answer = find_max(answer, start.Val + end.Val)
   return answer
}