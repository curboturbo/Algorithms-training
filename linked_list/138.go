package main



type Node2 struct {
    Val int
    Next *Node2
    Random *Node2
}

// solve with hash-table
func copyRandomList(head *Node2) *Node2 {
	translation:= map[*Node2]*Node2{}
    new_head := &Node2{Val:0, Next:nil, Random:nil}
	pointer:=new_head
	pointer2:=new_head
	for head != nil{
		new_node := &Node2{Val:head.Val, Next:nil, Random:head.Random}
		translation[head] = new_node
		new_head.Next=new_node
		new_head = new_head.Next
		head=head.Next
	}
	pointer=pointer.Next
	for pointer!= nil{
		link:=translation[new_head.Random]
		new_head.Random=link
		new_head = new_head.Next
	}
	return pointer2.Next
	
}

// solve using double copy and increasing first list
func copyRandomList2(head *Node2) *Node2 {
	if head == nil {
		return nil
	}

	// ============================
	// ШАГ 1: вставляем clone после каждой оригинальной ноды
	// Было: A -> B -> C
	// Стало: A -> A' -> B -> B' -> C -> C'
	// ============================
	curr := head
	for curr != nil {
		clone := &Node2{Val: curr.Val}

		clone.Next = curr.Next // clone указывает туда, куда раньше указывал curr
		curr.Next = clone      // curr теперь указывает на свой clone

		curr = clone.Next // переходим к следующей оригинальной ноде
	}

	// ============================
	// ШАГ 2: настраиваем Random у clone
	// curr.Random указывает на оригинал X
	// X.Next — это clone(X)
	// ============================
	curr = head
	for curr != nil {
		if curr.Random != nil {
			// curr.Next — это clone(curr)
			// curr.Random.Next — это clone(curr.Random)
			curr.Next.Random = curr.Random.Next
		}
		// перескакиваем через clone
		curr = curr.Next.Next
	}

	// ============================
	// ШАГ 3: разделяем оригинальный список и список копий
	// ============================
	curr = head
	cloneHead := head.Next // голова скопированного списка

	for curr != nil {
		clone := curr.Next

		// восстанавливаем оригинальный список
		curr.Next = clone.Next

		// связываем clone между собой
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}

		curr = curr.Next
	}

	return cloneHead
}

