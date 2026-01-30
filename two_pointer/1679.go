package main


func maxOperations(nums []int, k int) int {
    mapa:=map[int]int{}
	operations:=0
	for i:=0;i<len(nums);i++{
		_,request:= mapa[nums[i]]
		if request{
			mapa[nums[i]]++
		}else{mapa[nums[i]]=1}
	}
	for key,value := range mapa{
		target := k - key
		_,request:=mapa[target]
		if request{
			if key == target{
				operations+= value/2
			}else{
			operations+= min(mapa[target],value)
			delete(mapa,key)
			delete(mapa,target)
			}
		}else{
			delete(mapa,key)
		}
	}
	return operations
}

