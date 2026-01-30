package main


func max(x int,y int) int{
	if x>=y{return x}else{return y}
}

func lengthOfLongestSubstring(s string) int {
	a:=[]rune(s)
	mapa:=map[rune]int{}
	left:=0
	answer:=-1
	for i:=0;i<len(a);i++{
		_, ok := mapa[a[i]]
		if ok==false{
			mapa[a[i]] = 1
		}else{
			mapa[a[i]] +=1
			for mapa[a[i]] > 1 {
				mapa[a[left]]-=1
				left+=1
			}
		}
		answer = max(answer,i-left+1)
	}
	return answer
}
