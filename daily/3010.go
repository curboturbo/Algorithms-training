package main

func m(x int,y int)int{if x<=y{return x}else{return y}}


func minimumCost1(nums []int) int {
	if len(nums) == 3{return nums[0]+nums[1]+nums[2]}
	sum:=nums[0]
	min:=nums[1]
	cnt:=nums[1]
	for i:=2;i<len(nums);i++{
		cnt = m(cnt,min+nums[i])
		min = m(min,nums[i])
	}
	return sum+cnt
}
