package tree

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) PrePrintRecursion() {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	t.Left.PrePrintRecursion()
	t.Right.PrePrintRecursion()
	return
}

func (t *TreeNode) MidPrintRecursion() {
	if t == nil {
		return
	}
	t.Left.MidPrintRecursion()
	fmt.Println(t.Val)
	t.Right.MidPrintRecursion()
	return
}

func (t *TreeNode) PostPrintRecursion() {
	if t == nil {
		return
	}
	t.Left.PostPrintRecursion()
	t.Right.PostPrintRecursion()
	fmt.Println(t.Val)
	return
}

func (t *TreeNode) PrePrint() {
	cur := t
	tempList := list.New()
	for cur != nil || tempList.Len() != 0 {
		for cur != nil {
			fmt.Println(cur.Val)
			tempList.PushBack(cur)
			cur = cur.Left
		}
		if tempList.Len() != 0 {
			cur = tempList.Back().Value.(*TreeNode)
			tempList.Remove(tempList.Back())
			cur = cur.Right
		}
	}
}

func (t *TreeNode) MidPrint() {
	cur := t
	tempList := list.New()
	for cur != nil || tempList.Len() != 0 {
		for cur != nil {
			tempList.PushBack(cur)
			cur = cur.Left
		}
		if tempList.Len() != 0 {
			cur = tempList.Back().Value.(*TreeNode)
			tempList.Remove(tempList.Back())
			fmt.Println(cur.Val)
			cur = cur.Right
		}
	}
}

func (t *TreeNode) PostPrint() {
	if t == nil {
		return
	}
	cur := t
	tempList := list.New()
	tempList.PushBack(cur)
	var LastPos *TreeNode
	for tempList.Len() != 0 {
		cur = tempList.Back().Value.(*TreeNode)
		//1.左右节点均为空；2.右节点为空，左节点已经访问过;3.右节点已经访问过 => 则可以访问当前节点
		if (cur.Left == nil && cur.Right == nil) || (cur.Right == nil && LastPos == cur.Left) || LastPos == cur.Right {
			fmt.Println(cur.Val)
			LastPos = cur
			tempList.Remove(tempList.Back())
			continue
		}

		if cur.Right != nil {
			tempList.PushBack(cur.Right)
		}
		if cur.Left != nil {
			tempList.PushBack(cur.Left)
		}
	}
}

/*根据先序遍历和中序遍历重构二叉树，没有检错...*/
func ConstructTreeByPreAndMid(preNums []int, midNums []int) (root *TreeNode) {
	length := len(preNums)
	if length == 0 {
		return nil
	}
	root = &TreeNode{
		Val: preNums[0],
	}
	midInd := -1
	for i, v := range midNums {
		if v == preNums[0] {
			midInd = i
			break
		}
		if i == len(midNums) {
			return nil
		}
	}

	var lPre []int
	var lMid []int
	var rPre []int
	var rMid []int
	emptyNums := make([]int, 0)
	if midInd == 0 {
		lPre = emptyNums
		lMid = emptyNums
	} else {
		lPre = preNums[1 : 1+midInd]
		lMid = midNums[0:midInd]
	}

	if length == midInd+1 {
		rPre = emptyNums
		rMid = emptyNums
	} else {
		rPre = preNums[midInd+1:]
		rMid = midNums[midInd+1:]
	}

	root.Left = ConstructTreeByPreAndMid(lPre, lMid)
	root.Right = ConstructTreeByPreAndMid(rPre, rMid)

	return root
}

/*问题*/
/*
给定一个数组，求其以顺序存储方式构成的完全二叉树
*/
/*思路*/
/*
每一层的每个节点都依次设置左右节点
每设置一个节点，将该节点存入list
从上往下，从左往右设置节点
*/
func CoustructCompleteTree(nums []int) (root *TreeNode) {
	length := len(nums)
	if length == 0 {
		return nil
	}
	root = &TreeNode{
		Val: nums[0],
	}

	tempList := list.New()
	tempList.PushBack(root)
	for i := 1; i <= length/2; i++ {
		curPos := tempList.Front().Value.(*TreeNode)
		if (i*2 - 1) < length {
			left := &TreeNode{
				Val: nums[2*i-1],
			}
			curPos.Left = left
			tempList.PushBack(left)
		}
		if (i * 2) < length {
			right := &TreeNode{
				Val: nums[2*i],
			}
			curPos.Right = right
			tempList.PushBack(right)
		}
		tempList.Remove(tempList.Front())
	}

	return root
}

/*问题*/
/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

示例 1:

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
注意: 合并必须从两个树的根节点开始。
*/
/*思路*/
/*
递归，每次迭代函数只处理当前的val
其left和right交给下一次迭代函数计算得到
*/
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	var pl1 *TreeNode = nil
	var pl2 *TreeNode = nil
	var pr1 *TreeNode = nil
	var pr2 *TreeNode = nil
	val1 := 0
	val2 := 0
	if t1 != nil {
		pl1 = t1.Left
		pr1 = t1.Right
		val1 = t1.Val
	}
	if t2 != nil {
		pl2 = t2.Left
		pr2 = t2.Right
		val2 = t2.Val
	}
	return &TreeNode{
		Val:   val1 + val2,
		Left:  MergeTrees(pl1, pl2),
		Right: MergeTrees(pr1, pr2),
	}
}

/*问题*/
/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
/*思路*/
/*
递归大法好，当前层深度 = (max(1 + left深度,1 + right深度) || 0)
*/
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftDeep := MaxDepth(root.Left)
	rightDeep := MaxDepth(root.Right)
	res := leftDeep
	if rightDeep > leftDeep {
		res = rightDeep
	}
	return res + 1
}

/*问题*/
/*
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/
/*思路*/
/*
展开方式是先序
先先序遍历把结果存入数组，同时清空原数组
再重新生成二叉树
*/
func midLoopAndSetNil(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	midArr := make([]int, 1)
	midArr[0] = root.Val
	temp1 := midLoopAndSetNil(root.Left)
	midArr = append(midArr, temp1...)
	temp2 := midLoopAndSetNil(root.Right)
	midArr = append(midArr, temp2...)
	root.Left = nil
	root.Right = nil
	return midArr
}
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	midArr := midLoopAndSetNil(root)

	//此时root指针只存在当前值
	var p *TreeNode = nil
	for i := 1; i < len(midArr); i++ {
		if i == 1 {
			p = &TreeNode{
				Val: midArr[i],
			}
			root.Right = p
		} else {
			p.Right = &TreeNode{
				Val: midArr[i],
			}
			p = p.Right
		}
	}
	return
}

//func midLoopAndSetNil(root *TreeNode, midArr []int) {
//	if root == nil {
//		return
//	}
//	midArr = append(midArr, root.Val)
//	midLoopAndSetNil(root.Left, midArr)
//	midLoopAndSetNil(root.Right, midArr)
//	root.Left = nil
//	root.Right = nil
//	return
//}
//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	midArr := make([]int, 0)
//	midLoopAndSetNil(root, midArr)
//	fmt.Println(midArr)
//
//	//此时root指针只存在当前值
//	var p *TreeNode = nil
//	for i := 1; i < len(midArr); i++ {
//		if i == 1 {
//			p = &TreeNode{
//				Val: midArr[i],
//			}
//			root.Right = p
//		} else {
//			p.Right = &TreeNode{
//				Val: midArr[i],
//			}
//			p = p.Right
//		}
//	}
//	return
//}
