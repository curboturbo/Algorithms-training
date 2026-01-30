package main
import (
	"fmt"
	"sort"
	"math"
)
func m(x int, y int) int{if x<=y{return x}else{return y}}

func coinChange(coins []int, amount int) int {
	dp:= make([]int, amount+1)
	for i:=1;i<amount+1;i++{
		dp[i] = math.MaxInt32
	}
	//cnt:=0
	dp[0] = 0
	sort.Slice(coins,func(i, j int) bool {return coins[i]<coins[j]})
	for i:=1;i<amount+1;i++{
		for j:=0;j<len(coins);j++{
			if i - coins[j] >=0{
				dp[i] = m(dp[i-coins[j]]+1,dp[i])
			}
		}
	}
	for i:=0;i<len(dp);i++{
		fmt.Printf("%d ",dp[i])
	}
	if dp[amount] == math.MaxInt32{return -1}else{return dp[amount]}
}

func main(){
	a:=[]int{1,2,5}
	coinChange(a,11)
}