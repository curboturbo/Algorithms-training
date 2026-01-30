package main

//[2,3,1,2,4,3]


func min(x int,y int) int{
	if x<=y{return x}else{return y}
}

func minSubArrayLen(target int, nums []int) int {
    total:=0
	left:=0
	flag := false
	ans := len(nums)+1
	for i:=0;i<len(nums);i++{
		total += nums[i]
		for total > target{
			if total >= target{
			ans = min(ans,i-left+1)
			flag = true
			}
			if left==i{break}
			total-=nums[left]
			left+=1
		}
		if total == target{
			ans = min(ans,i-left+1)
			flag = true
		}
	}
	if flag{return ans}else{return 0}
}