package main
import "math"

func dfs(root *TreeNode, maxSoFar int, count *int) {
	if root == nil {
		return
	}
	if root.Val >= maxSoFar {
		*count++
		maxSoFar = root.Val
	}
	dfs(root.Left, maxSoFar, count)
	dfs(root.Right, maxSoFar, count)
}

func goodNodes(root *TreeNode) int {
	count := 0
	dfs(root, math.MinInt32, &count)
	return count
}
