package main



func BinarySearch(arr []int) int {
	left, right := 0, len(arr)-1
	firstNeg := len(arr)

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] < 0 {
			firstNeg = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return len(arr) - firstNeg
}

func countNegatives(grid [][]int) int {
	cnt:=0
    for i:=0;i<len(grid);i++{
		cnt+=BinarySearch(grid[i])
	}
	return cnt
}


//0-14
//[9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5]