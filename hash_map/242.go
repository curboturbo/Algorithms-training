package main

import "fmt"

func isAnagram(pattern string, s string) bool {
    ransom:=[]rune(pattern)
	mag:= []rune(s)
	mapa:=map[rune]int{}
	for i:=0;i<len(ransom);i++{
		_, response:= mapa[ransom[i]]
		if response{
			mapa[ransom[i]]+=1
		}else{mapa[ransom[i]]=1}	
	}
	for i:=0;i<len(mag);i++{
		_, response:= mapa[mag[i]]
		if response{mapa[mag[i]]-=1}else{continue}
	}
	for _, value:= range mapa{
		if value ==0{continue}else{return false}
	}
	return true
}

func main(){
	fmt.Print(wordPattern("cat","rat"))
}