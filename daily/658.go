package main

import (
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if x <= arr[0] {
		return arr[:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	left := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= x && arr[i+1] > x {
			left = i
			break
		}
	}
	right := left + 1
	answer := []int{}
	var a int
	var b int
	for k > 0 {
		if left >= 0 {
			a = arr[left]
		} else {
			a = -1000000 
		}
		if right < len(arr) {
			b = arr[right]
		} else {
			b = 1000000
		}
		if left >= 0 && (right >= len(arr) || abs(x-a) <= abs(x-b)) {
			answer = append(answer, a)
			left--
		} else {
			answer = append(answer, b)
			right++
		}
		k--
	}
	sort.Ints(answer)
	return answer
}

