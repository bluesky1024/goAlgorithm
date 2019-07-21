package tree

import (
	"fmt"
	"testing"
)

func TestTreeNode_Print(t *testing.T) {
	nums := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	root := ConstructCompleteTree(nums)
	root.MidPrint()
}

func TestConstructTreeByPreAndMid(t *testing.T) {
	//preNums := []int{1, 2, 4, 5, 3, 6}
	//midNums := []int{4, 2, 5, 1, 3, 6}
	preNums := []int{1, 2, 3, 4}
	midNums := []int{1, 2, 3, 4}
	root := ConstructTreeByPreAndMid(preNums, midNums)
	root.PrePrint()
	fmt.Println("---")
	root.MidPrint()
}

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
	root := ConstructCompleteTree(nums)
	root.MidPrint()
}

func TestCoustructTreeInLevel(t *testing.T) {
	nums := []int{10, 5, 15, 3, 7, -1, 18, 1, -1, 6, -1, -1, -1, 21, -1}
	root := ConstructTreeInLevel(nums)
	fmt.Println("mid print")
	root.MidPrint()
	fmt.Println("pre print")
	root.PrePrint()
}

func TestMaxAncestorDiff(t *testing.T) {
	root := ConstructTreeInLevel([]int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13})
	fmt.Println(MaxAncestorDiff(root))
}

func TestSumNumbers(t *testing.T) {
	root := ConstructTreeInLevel([]int{1, 2, 3})
	fmt.Println(SumNumbers(root))
}
