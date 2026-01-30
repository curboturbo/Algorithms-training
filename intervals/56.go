package main
import "sort"


func max(x int, y int) int{
    if x>=y{return x}else{return y}
}


func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
	})
	begin:=intervals[0][0]
	end:=intervals[0][1]
	array:=[][]int{}
	
	for i:=1;i<len(intervals);i++{
		new_begin:=intervals[i][0]
		new_end:=intervals[i][1]
		if (end==new_end) && (begin==new_begin){
			end = new_end
			begin = new_begin
		}else if new_begin <=end{
			end = max(end,new_end)
		}else{
			array = append(array,[]int{begin,end})
			begin = new_begin
			end = new_end
		}
	}
	array = append(array, []int{begin,end})
	return array
}