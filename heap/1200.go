package main


import "sort"

func m(x int, y int)int{if x<=y{return x}else{return y}}

func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]})
	a:=arr[0]
	b:=arr[1]
	min:=m(a,b)
	ans := ([][]int{})
	for i:=1;i<len(arr);i++{
		dist:=arr[i]-arr[i-1]
		if dist < min{
			ans = nil
			min = dist
			//ans := new([][]int{})
			ans = append(ans, []int{arr[i],arr[i-1]})
		}else if dist == min{
			ans = append(ans, []int{arr[i],arr[i-1]})
		}else{continue}
	}
	return ans
}
func main(){
	a:=[]int{4,2,1,3}
	minimumAbsDifference(a)
}