package main
import "fmt"


func Max(x int, y int) int{
	if x>=y{return x}else{return y}
}

func longestConsecutive(nums []int) int {
    mapa := map[int]bool{}
    for _, n := range nums {
        mapa[n] = true
	}
    maxSum := 0
    for n := range mapa {
        if !mapa[n-1] { 
            cnt := 1
            for mapa[n+cnt] {
                cnt++
            }
            if cnt > maxSum {
                maxSum = cnt
            }
        }
    }
    return maxSum
}

func main(){
	fmt.Print(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}