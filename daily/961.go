package main
func repeatedNTimes(nums []int) int {
    mapa:=map[int]int{}
	cnt:=0
	for i:=0;i<len(nums);i++{
		_,response:=mapa[nums[i]]
		if response{
			mapa[nums[i]]+=1
			if mapa[nums[i]] == 2{cnt--}
		}else{
			mapa[nums[i]]=1
			cnt++
		}
	}
	//cnt-  кол-ов уникальных
	for i:=0;i<len(nums);i++{
		if mapa[nums[i]] == cnt{return nums[i]}
	}
	return 0

}