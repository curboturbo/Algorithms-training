package main


func finder(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*slice = append(*slice, root.Val)
		return
	}
	finder(root.Left, slice)
	finder(root.Right, slice)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	first_order := []int{}
	second_order := []int{}
	finder(root1, &first_order)
	finder(root2, &second_order)
	if len(first_order) != len(second_order) {
		return false
	}
	for i := range first_order {
		if first_order[i] != second_order[i] {
			return false
		}
	}
	return true
}
