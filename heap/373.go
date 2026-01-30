package main
type Node struct{
	I int
	J int 
	Sum int
}


func push1(heap *[]Node, el *Node){
	(*heap) = append((*heap), *el)
	pos:=len(*heap)-1
	balance:=false
	for !balance{
		parent:= (pos-1)/2
		if pos - 1 < 0{break}
		if (*heap)[parent].Sum < (*heap)[pos].Sum{balance=true}else{
			(*heap)[parent],(*heap)[pos] = (*heap)[pos],(*heap)[parent]
			pos = parent
		}
	}
}

func pop1(heap *[]Node) Node {
    ans := (*heap)[0]
    last := (*heap)[len(*heap)-1]
    *heap = (*heap)[:len(*heap)-1]
    if len(*heap) == 0 {
        return ans
    }
    (*heap)[0] = last
    pos := 0
    n := len(*heap)
    balance := false
    for !balance {
        left := 2*pos + 1
        right := 2*pos + 2
        smallest := pos
        if left < n && (*heap)[left].Sum < (*heap)[smallest].Sum {
            smallest = left
        }
        if right < n && (*heap)[right].Sum < (*heap)[smallest].Sum {
            smallest = right
        }
        if pos == smallest {
            balance = true
        } else {
            (*heap)[pos], (*heap)[smallest] = (*heap)[smallest], (*heap)[pos]
            pos = smallest
        }
    }
    return ans
}

func m1(x int, y int) int{if x<=y{return x}else {return y}}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    heap := []Node{}
    answer := [][]int{}
    border := m1(len(nums2), k) // сколько элементов из первой строки кладём

    // Инициализация: первая строка (nums1[0]) или ограничение по k
    for j := 0; j < border; j++ {
        push1(&heap, &Node{I: 0, J: j, Sum: nums1[0] + nums2[j]})
    }

    for k > 0 && len(heap) > 0 {
        top := pop1(&heap)
        answer = append(answer, []int{nums1[top.I], nums2[top.J]})
        k--

        // добавляем следующий элемент в столбце (следующий по i)
        if top.I+1 < len(nums1) {
            push1(&heap, &Node{
                I:   top.I + 1,
                J:   top.J,
                Sum: nums1[top.I+1] + nums2[top.J],
            })
        }
    }
    return answer
}
