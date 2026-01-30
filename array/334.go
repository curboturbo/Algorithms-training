package main
import "fmt"


//Input: nums = [1,2,3,4,5]
//Output: true
//Explanation: Any triplet where i < j < k is valid.
//Example 2:
//
//Input: nums = [5,4,3,2,1]
//Output: false
//Explanation: No triplet exists.
//Example 3:
//
//Input: nums = [2,1,5,0,4,6]
//Output: true
//Explanation: One of the valid triplet is (1, 4, 5), because nums[1] == 1 < nums[4] == 4 < nums[5] == 6.


func min(x int, y int) int {if x<=y{return y}else{return x}}

func increasingTriplet(nums []int) bool {
	if len(nums) <= 2{return false}
	most_min:=1234112
	next_min:=12312
	for i:=0;i<len(nums);i++{
		if nums[i] <= most_min{
			most_min = nums[i]
		}else if nums[i]<=next_min{
			next_min=nums[i]
		}else{return true}
	}
	return false   
}


func main(){
	a:=[]int{20,100,10,12,5,13}
	fmt.Print(increasingTriplet(a))
}