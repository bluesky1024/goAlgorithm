package tree

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 5,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 6,
	}

	Flatten(root)
}

func TestCoustructCompleteTree(t *testing.T) {
	nums := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	root := CoustructCompleteTree(nums)
	root.MidPrint()
}
