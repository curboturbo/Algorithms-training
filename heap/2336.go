package main



///КОШМАР А НЕ ЗАДАЧА КУЧИ УБИВАЮТ

type SmallestInfiniteSet struct {
	Cnt int // указатель на следующий минимальный
	Heap *[]int // ссылка на кучу
	Mapa map[int]int
}


func Constructor() SmallestInfiniteSet {
	heap := new([]int)
	mapa:=map[int]int{}
	object:=&SmallestInfiniteSet{Cnt:1, Heap: heap,Mapa:mapa}
	return *object
}


func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(*this.Heap) > 0 {
		return pop(this.Heap,this)
	}
	res := this.Cnt
	this.Cnt++
	return res
}

// протащили вверх минмальный
func push(heap *[]int, num int){
	(*heap) = append((*heap), num)
	pos:=len(*heap)-1
	balance:=false
	for !balance{
		parent:= (pos-1)/2
		if pos - 1 < 0{break}
		if (*heap)[parent] < (*heap)[pos]{balance=true}else{
			(*heap)[parent],(*heap)[pos] = (*heap)[pos],(*heap)[parent]
			pos = parent
		}
	}
}


func pop(heap *[]int,this *SmallestInfiniteSet) int{
	ans:=(*heap)[0] // return ans
	_,request:=this.Mapa[ans]
	if request{delete(this.Mapa, ans)}
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



func (this *SmallestInfiniteSet) AddBack(num int){
	if num < this.Cnt{
		this.Mapa[num]++
		if this.Mapa[num]==1{
			push(this.Heap, num)
		}
	}
}