package main
import "sort"
import "fmt"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	answer := [][]int{}
	seen := map[string]bool{} // аналог множества сишного

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s == 0 {
				triplet := fmt.Sprintf("%d,%d,%d", nums[i], nums[left], nums[right])
				if _, found := seen[triplet]; !found {
					answer = append(answer, []int{nums[i], nums[left], nums[right]})
					seen[triplet] = true
				}
				left++
				right--
			} else if s < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return answer
}
func main(){
	c:= threeSum([]int{4,-1,-1,0,1,2})
	fmt.Print(c)
}