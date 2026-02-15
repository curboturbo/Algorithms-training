package main



type Trie struct {
	Terminal bool
	Value rune
	Nodes [26]*Trie
}


func Constructor() Trie{
	object:=&Trie{Terminal: false, Nodes:[26]*Trie{}, Value:0}
	return *object   
}


func (this *Trie) Insert(word string) {
    if len(word) == 0{
        this.Terminal = true
        return
    }
    ch := rune(word[0])
    index := ch - 'a'
    if this.Nodes[index] == nil {
        this.Nodes[index] = &Trie{Value: ch}
    }
    this.Nodes[index].Insert(word[1:])
}


func (this *Trie) Search(word string) bool {
	if len(word) == 0 {return this.Terminal}
	ch:=rune(word[0])
	index:=ch - 'a'
	if this.Nodes[index] == nil{
		return false
	}
	return this.Nodes[index].Search(word[1:])
}


func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0{return true}
	ch:=rune(prefix[0])
	index:=ch - 'a'
	if this.Nodes[index] == nil{
		return false
	}
	return this.Nodes[index].StartsWith(prefix[1:])
}