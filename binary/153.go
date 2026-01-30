package main




//[4,5,6,7,0,1,2]
func findMin(nums []int) int {
	left:=0
	right:=len(nums)-1
	if nums[left] < nums[right]{return nums[left]}
	var mid int
	for left < right{
		mid = (left+right)/2
		if nums[mid] > nums[right] {
    		left = mid + 1
		} else {
    		right = mid
		}
	}
	return nums[left]
}