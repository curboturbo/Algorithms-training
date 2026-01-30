package main
import "fmt"

func m(x int,y int) int{if x<=y{return x}else{return y}}


func minCostClimbingStairs(cost []int) int {
    cost = append(cost,0)
    cost = append(cost,0)
    n := len(cost)
    dp:=make([]int,n)
    dp[0] = cost[0]
    dp[1] = cost[1]
    for i:=2;i<len(cost);i++{
        dp[i] = cost[i]+m(dp[i-1],dp[i-2])
		fmt.Printf("%d ", dp[i])
    }
	return m(dp[n-1],dp[n-2])
}
func main(){
	minCostClimbingStairs([]int{10,15,20})
}