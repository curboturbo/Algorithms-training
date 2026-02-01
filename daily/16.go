package main
import (
	"sort"
)
func mod(a int) int {
	if a < 0 {return -a}
	return a
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {continue}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {return target}
			if mod(sum-target) < mod(ans-target) {
				ans = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}