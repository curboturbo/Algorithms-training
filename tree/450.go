package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{return root}
	if key > root.Val{
		root.Right = deleteNode(root.Right,key)
	}else if key < root.Val{
		root.Left = deleteNode(root.Left, key)
	}else{
		// то есть мы нашли удаляемую вершину
		if root.Left!=nil && root.Right == nil{
			return root.Left
		}else if root.Left==nil && root.Right!=nil{
			return root.Right
		}else if root.Left==nil && root.Right == nil{return nil
		}else{
			rec:= root.Right
			for rec.Left != nil{
				rec=rec.Left
			}
			root.Val = rec.Val
			root.Right = deleteNode(root.Right,rec.Val)
		}
	}
	return root
}