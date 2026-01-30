package main


func max(x int,y int) int{
	if x>=y{
		return x
	}else{return y}
}
func min(x int,y int) int{
	if x<=y{return x}else{return y}
}

func maxArea(height []int) int {
	S:=0
	left:=0
	right:=len(height)-1
	for left < right{
		S = max(S,min(height[left],height[right])*(right-left))
		if height[left]<=height[right]{
			left +=1
		}else{
			right-=1
		}
	}
	return S
    
}

