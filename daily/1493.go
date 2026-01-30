package main


func Max(x int,y int) int{
	if x>=y{return x}else{return y}
}


func longestSubarray(nums []int) int {
	death_count:=0
	max:=-1
	left_pointer := 0
	cur:=0
	for i:=0;i<len(nums);i++{
		if nums[i] == 0{
			death_count++
		}
		cur++
		if death_count == 2{
			max = Max(max,cur-1)
			for death_count > 1{
				if nums[left_pointer] == 0{death_count--}
				left_pointer++
				cur--
			}
		}
		max = Max(max,cur)
	}
	return max-1
}