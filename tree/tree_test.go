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

func TestConstrucTreeInLevelWithoutInvalidNode(t *testing.T) {
	nums := []int{9, 6, -3, -1, -1, -6, 2, -1, -1, 2, -1, -6, -6, -6}
	root := ConstrucTreeInLevelWithoutInvalidNode(nums)
	root.PrePrint()
	fmt.Println("---")
	root.MidPrint()
}

func TestConstructTreeByPreAndMid(t *testing.T) {
	root := &TreeNode{
		Val: 9,
	}
	root.Left = &TreeNode{
		Val: 6,
	}
	root.Right = &TreeNode{
		Val: -3,
		Left: &TreeNode{
			Val: -6,
		},
	}

	temp22 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: -6,
			Left: &TreeNode{
				Val: -6,
			},
		},
		Right: &TreeNode{
			Val: -6,
		},
	}
	temp2 := &TreeNode{
		Val:  2,
		Left: temp22,
	}
	root.Right.Right = temp2

	//root.PrePrint()
	//fmt.Println("---")
	//root.MidPrint()

	a := ConstructTreeByPreAndMid([]int{9, 6, -3, -6, 2, 2, -6, -6, -6}, []int{6, 9, -6, -3, -6, -6, 2, -6, 2})
	//fmt.Println(root, a)
	a.PrePrint()
	a.MidPrint()

	////preNums := []int{1, 2, 4, 5, 3, 6}
	////midNums := []int{4, 2, 5, 1, 3, 6}
	//preNums := []int{1, 2, 3, 4}
	//midNums := []int{1, 2, 3, 4}
	//root := ConstructTreeByPreAndMid(preNums, midNums)
	//root.PrePrint()
	//fmt.Println("---")
	//root.MidPrint()
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

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 9,
	}
	root.Left = &TreeNode{
		Val: 6,
	}
	root.Right = &TreeNode{
		Val: -3,
		Left: &TreeNode{
			Val: -6,
		},
	}

	temp22 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: -6,
			Left: &TreeNode{
				Val: -6,
			},
		},
		Right: &TreeNode{
			Val: -6,
		},
	}
	temp2 := &TreeNode{
		Val:  2,
		Left: temp22,
	}
	root.Right.Right = temp2
	fmt.Println(MaxPathSum(root))
}

func TestConstructNewTreeWithParent(t *testing.T) {
	nums := []int{9, 6, -3, -1, -1, -6, 2, -1, -1, 2, -1, -6, -6, -6}
	root := ConstrucTreeInLevelWithoutInvalidNode(nums)
	newRoot := ConstructNewTreeWithParent(root, nil, false)
	fmt.Println(newRoot)
}

func TestGetNodesByPre(t *testing.T) {
	nums := []int{9, 6, -3, -1, -1, -6, 2, -1, -1, 2, -1, -6, -6, -6}
	root := ConstrucTreeInLevelWithoutInvalidNode(nums)
	list := GetNodesByPre(root)
	for _, v := range list {
		fmt.Println(v.Val)
	}
	fmt.Println("---")
	root.PrePrint()
}

func TestFindDuplicateSubtrees(t *testing.T) {
	nums := []int{0, 0, 0, 0, -1, -1, 0, -1, -1, -1, 0}
	root := ConstrucTreeInLevelWithoutInvalidNode(nums)

	//str := GetLevelLoopString([]*TreeNode{root})
	//fmt.Println(str)

	list := FindDuplicateSubtrees(root)
	fmt.Println(list)
	//for _, v := range list {
	//	fmt.Println(v.Val)
	//}
}
