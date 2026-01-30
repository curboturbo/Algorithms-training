package main


func bin(nums []int, target int,left_flag bool) int{
	left:=0
	right:=len(nums)-1
	var mid int
	var idx int = -1
	for left <= right{
		mid = (left+right)/2
		if nums[mid] > target{
			right = mid - 1
		}else if target > nums[mid]{
			left = mid + 1
		}else if target == nums[mid]{
			idx = mid
			if left_flag{
				right = mid-1
			}else{left = mid+1}
		}
	}
	return idx
}


func searchRange(nums []int, target int) []int {
    begin := bin(nums, target, true)
    end := bin(nums, target, false)
    return []int{begin, end}
}
