package main

// using hash-table
func hasCycle(head *ListNode) bool {
    mapa:=map[*ListNode]int{}
	for head!=nil{
		_,response:= mapa[head]
		if !response{
			mapa[head] = 1
		}else{
			return true
		}
		head=head.Next
	}
	return false
}
//using tow pointers
func hasCycle2(head *ListNode) bool{
	var turtle *ListNode = head
	var rabbit *ListNode = head
	for rabbit!=nil && rabbit.Next!=nil{
		turtle = turtle.Next
		rabbit = rabbit.Next.Next
		if turtle==rabbit{return true}
		
	}
	return false
}
