package main
import "fmt"

//[1,3,5,6], target = 5
func searchInsert(nums []int, target int) int {
	pointer:=new(int)
	var mid int
	left:=0
	right:=len(nums)-1
	for left <= right{
		mid = (left+ right)/2
		if nums[mid] == target{
			*pointer = mid
			return mid
		}
		if target < nums[mid]{
			right = mid-1
		}else if (target > nums[mid]){
			left = mid + 1
		}
	}
	return left
}


func main(){
	a:=[]int{1,3,5,6}
	fmt.Print(searchInsert(a,2))

}