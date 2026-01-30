package main

//import "fmt"

// head and tail this is virtaul elements, cant be filled by other elem, this is just a pointers on start and end element
type List struct {
	Key int
	Value int
	Next *List
	Previous *List
}


type LRUCache struct {
	Capacity int
	Head *List
	Tail *List
	Size int
	Mapa map[int]*List
}


func Constructor(capacity int) LRUCache {
	head := &List{}
	tail := &List{}
	head.Next = tail
	head.Previous = nil
	tail.Previous = head
	tail.Next = nil
	mapa := map[int]*List{}
	return LRUCache{
		Capacity: capacity,
		Head:head,
		Tail:tail,
		Mapa:mapa,
		Size:0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.Mapa[key]
	if !exists {
		return -1
	}
	this.removeNode(node)
	this.addToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.Mapa[key]; exists {
		node.Value = value
		this.removeNode(node)
		this.addToHead(node)
		return
	}

	newNode := &List{Key: key, Value: value}

	if this.Size < this.Capacity {
		this.addToHead(newNode)
		this.Mapa[key] = newNode
		this.Size++
	} else {
		lru := this.Tail.Previous
		this.removeNode(lru)
		delete(this.Mapa, lru.Key)
		this.addToHead(newNode)
		this.Mapa[key] = newNode
	}
}

func (this *LRUCache) removeNode(node *List) {
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
}

func (this *LRUCache) addToHead(node *List) {
	node.Next = this.Head.Next
	node.Previous = this.Head
	this.Head.Next.Previous = node
	this.Head.Next = node
}


// main() {
//	cache := Constructor(2)
//	fmt.Println("Put(1,1)")
//	cache.Put(1, 1)
//	fmt.Println("Put(2,2)")
//	cache.Put(2, 2)
//	fmt.Println("Get(1):", cache.Get(1)) // ожидаем 1
//	fmt.Println("Put(3,3) (должен вытеснить 2)")
//	cache.Put(3, 3)
//	fmt.Println("Get(2):", cache.Get(2)) // ожидаем -1
//	fmt.Println("Put(4,4) (должен вытеснить 1)")
//	cache.Put(4, 4)
//	fmt.Println("Get(1):", cache.Get(1)) // -1
//	fmt.Println("Get(3):", cache.Get(3)) // 3
//	fmt.Println("Get(4):", cache.Get(4)) // 4
//}
