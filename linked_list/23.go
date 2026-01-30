package main



func pop(heap *[]int) int{
	ans:=(*heap)[0] // return ans
	last := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	if len(*heap) == 0{return ans}
	(*heap)[0] = last
	pos:=0
	n:= len(*heap)
	balance:=false
	for !balance{
		left:= (2*pos)+1
		right:=(2*pos)+2
		smallest:=pos
		if left < n && (*heap)[left] < (*heap)[smallest]{
			smallest = left
		}
		if right < n && (*heap)[right] < (*heap)[smallest]{
			smallest = right
		}
		if pos == smallest{balance = true}else{
			(*heap)[pos],(*heap)[smallest] = (*heap)[smallest],(*heap)[pos]
		}
		pos = smallest
	}
	return ans
}

func insert(heap *[]int, element int){
	(*heap) = append((*heap), element)
	position:=len((*heap))-1
	balance:=false
	var parent int
	for !balance{
		if position - 1 < 0{break}
		parent = (position-1)/2
		if (*heap)[parent] >= (*heap)[position]{
			(*heap)[parent],(*heap)[position] = (*heap)[position],(*heap)[parent]
			position = parent
		}else{break}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	new_root:=&ListNode{Val:0,Next:nil}
	pointer:=new_root
	heap:=new([]int)
	if len(lists) == 0 {return nil}
	for i:=0;i<len(lists);i++{
		root:= lists[i]
		for root != nil{
			insert(heap,root.Val)
			root = root.Next
		}
	}
	for len((*heap)) > 0 {
		top:= pop(heap)
		children:=&ListNode{Val:top,Next:nil}
		new_root.Next = children
		new_root = new_root.Next
	}
	return pointer.Next
}
