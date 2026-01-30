package main
import "fmt"
import "sort"


func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
    return points[i][1] < points[j][1]
	})
	flag:=true
	arrows:=0
	shot_position:=0
	for i:=1;i<len(points);i++{
		if points[shot_position][1] >= points[i][0]{
			flag= true
			continue
		}else{
			arrows++
			shot_position = i
			flag = false
		}
	}
	if flag{arrows++}
	return arrows
}

func main(){
	fmt.Print(findMinArrowShots([][]int{{10,16},{2,8},{1,6},{7,12}}))
}