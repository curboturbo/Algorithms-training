package main



// организация k-арной кучи
func balance(heap *[]int, element int){
	*heap = append(*heap,element)
	pos:=len(*heap)-1
	balance:=false
	for !balance{
		if (pos - 1) < 0{break}
		par:=(pos-1)/2
		if (*heap)[pos] >= (*heap)[par]{balance=true}else{
			temp:=(*heap)[par]
			(*heap)[par] = (*heap)[pos]
			(*heap)[pos] = temp
			pos = par
		}
	}
}


func balance_insert(heap *[]int, element int) {
	if element <= (*heap)[0] {
		return
	}
	(*heap)[0] = element
	n := len((*heap))-1
	pos:=0
	// "просеиваем вниз", чтобы куча оставалась min-кучей
	balance:= false
	for !balance{
		left:= (pos*2)+1
		right:=(pos*2)+2
		smallest:= pos
		if left < n && (*heap)[left] < (*heap)[smallest] {
			smallest = left
		}
		if right < n && (*heap)[right] < (*heap)[smallest]{
			smallest = right
		}
		if smallest == pos{balance=true}else{
			(*heap)[pos], (*heap)[smallest] = (*heap)[smallest], (*heap)[pos]
			pos = smallest
		}
	}
}


func findKthLargest(nums []int, k int) int {
    heap:= new([]int)
	for i:=0;i<k;i++{
		balance(heap, nums[i])
	}
	for i:=k;i<len(nums);i++{
		balance_insert(heap,nums[i])
	}
	return (*heap)[0]	
}