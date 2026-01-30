package main

func mi(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func path(root *TreeNode, ptr *int, prev int, cnt int) {
	if root == nil {
		return
	}
	*ptr = mi(*ptr, cnt)
	if prev == 1 {
		path(root.Left, ptr, 0, cnt+1)
		path(root.Right, ptr, 1, 1)
	} else {
		path(root.Right, ptr, 1, cnt+1)
		path(root.Left, ptr, 0, 1)
	}
}

func longestZigZag(root *TreeNode) int {
	ptr := new(int)
	path(root.Left, ptr, 0, 1)
	path(root.Right, ptr, 1, 1)
	return *ptr
}
