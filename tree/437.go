package main

func find(root *TreeNode, hashMap map[int]int, sum int, target int, ptr *int) {
	if root == nil {
		return
	}
	sum += root.Val
	*ptr += hashMap[sum-target]
	hashMap[sum]++
	find(root.Left, hashMap, sum, target, ptr)
	find(root.Right, hashMap, sum, target, ptr)
	hashMap[sum]--
}

func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	hashMap := map[int]int{0: 1}
	find(root, hashMap, 0, targetSum, &count)
	return count
}
