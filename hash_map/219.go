package main

import "fmt"
//Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
//Example 1:
//Input: nums = [1,2,3,1], k = 3
//Output: true
func containsNearbyDuplicate(nums []int, k int) bool {
	mapa:=map[int]int{}
	var cnt int = 0
	k+=1
	for i:=0;i<len(nums);i++{
		cnt += 1
		_,response:=mapa[nums[i]]
		if response{
			mapa[nums[i]]+=1
			return true
		}else{mapa[nums[i]]=1}
		if cnt==k{
			cnt-=1
			if mapa[nums[i-cnt]]>1{mapa[nums[i-cnt]]-=1}else{delete(mapa,nums[i-cnt])}
		}
	}
	return false
}
func main(){
	fmt.Print(containsNearbyDuplicate([]int{1,2,3,1},3))
}