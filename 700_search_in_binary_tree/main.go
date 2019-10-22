package main

// TreeNode basic format
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SearchBST to search a binary tree
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val < val:
		return SearchBST(root.Right, val)
	case val < root.Val:
		return SearchBST(root.Left, val)
	default:
		return root
	}
}
