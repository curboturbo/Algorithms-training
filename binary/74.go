package main

import  "fmt"
func searchMatrix(matrix [][]int, target int) bool {
    left:=0
	right:=len(matrix)-1
	m:=len(matrix[0])
	var mid int
	for left<=right{
		mid = (left+right)/2
		if target > matrix[mid][m-1]{
			left = mid+1
		}else if target < matrix[mid][m-1]{
			right = mid-1
		}else{break}
	}
	pointer := left
	left = 0
	right = m-1
	for left <= right{
		mid := (left+right)/2
		if target < matrix[pointer][mid]{
			right = mid-1
		}else if target > matrix[pointer][mid]{
			left = mid+1
		}else if (target == matrix[pointer][mid]){return true}
	}
	//fmt.Print(left)
	return false
}

func main(){
	a:=[][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60},}
	searchMatrix(a,3)
}