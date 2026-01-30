package main


import "fmt"
// 1 0 0 0 1 -> true
// 1 1 1 0 0
func canPlaceFlowers(flowerbed []int, n int) bool {
	prev:=0
	for i:=0;i<len(flowerbed);i++{
		if i == len(flowerbed)-1{
			if prev == 0 && flowerbed[i]==0{
				n--
				break
			}else{break}
		}

		if prev+flowerbed[i]+flowerbed[i+1] == 0{
			prev = 1
			n -- 
		}else{prev = flowerbed[i]}
	}
	if n == 0 {
    return true
	} else {
   		return false
	}
}

func main(){
	a:=[]int{1,0,0,0,1}
	fmt.Print(canPlaceFlowers(a,2))
}