package main
import "fmt"
func m(x int, y int)int{if x>=y{return x}else{return y}}


func lengthOfLIS(nums []int) int {
	dp:=make([]int, len(nums))
	for i:=0;i<len(nums);i++{dp[i] = 1}
	dp[0]=1
    mx:=-1000
	for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[j] < nums[i] && dp[i] < dp[j]+1{
				dp[i] = dp[j]+1
                mx = m(mx,dp[i])
			}
		}
	}
	for i:=0;i<len(nums);i++{fmt.Printf("%d ",dp[i])}
    if mx == -1000{return 1}else{return mx}
}