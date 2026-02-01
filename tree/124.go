/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main
import "math"
func mx(x int,y int) int{if x>=y{return x}else{return y}}

func dp(root *TreeNode,pointer *int) int{
	if root==nil{return 0}
	l:= mx(dp(root.Left,pointer),0)
	r:=mx(dp(root.Right,pointer),0)
	*pointer = mx(*pointer,root.Val + l + r)
	return root.Val + mx(l,r)
}
func maxPathSum(root *TreeNode) int {
    pointer:=new(int)
	*pointer = math.MinInt64
	dp(root,pointer)
	return *pointer
}