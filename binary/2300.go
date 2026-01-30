package main
import "sort"

//4

//> 28
// 5 5 7

func bin_except(nums []int, elem int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid]*elem < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return len(nums) - left
}


func successfulPairs(spells []int, potions []int, success int64) []int {
    a:=[]int{}
	sort.Slice(potions,func(i, j int) bool {return potions[i]<=potions[j]})
	for i:=0;i<len(spells);i++{
		result := bin_except(potions,spells[i],int(success))
		a = append(a, result)
	}
	return a
}