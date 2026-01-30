package main

import "sort"


//happiness = [2,3,4,5] k = 1


func maximumHappinessSum(happiness []int, k int) int64 {
	var decrease int = 0
	var totalSum int64 = 0
	sort.Ints(happiness)
	for i := len(happiness) - 1; i >= 0; i-- {
		if k > 0 {
			if happiness[i]-decrease > 0 {
				totalSum += int64(happiness[i] - decrease)
				decrease++
				k--
			}
		}
	}
	return totalSum
}

