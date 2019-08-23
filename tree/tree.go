package tree

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
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

func CompareTwoTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 == nil {
		return false
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return CompareTwoTree(root1.Left, root2.Right) && CompareTwoTree(root1.Right, root2.Right)
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
func ConstructCompleteTree(nums []int) (root *TreeNode) {
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
给定一个数组，[9 6 -3 nil nil -6 2 nil nil 2 nil -6 -6 -6]，属于二叉树的层次遍历，其中-1表示该节点为null
若某节点处为nil，则下一层遍历中，该节点处不再遍历
重构二叉树
*/
/*思路*/
/*
初始化一个
*/
func ConstructTreeInLevelWithoutInvalidNode(nums []int) (root *TreeNode) {
	length := len(nums)
	if length == 0 {
		return nil
	}
	root = &TreeNode{
		Val: nums[0],
	}
	var nodeLevel []*TreeNode
	nodeLevel = append(nodeLevel, root)
	ind := 1
	for ind < length {
		var tempNodes []*TreeNode
		for _, node := range nodeLevel {
			if ind >= length {
				break
			}
			if nums[ind] != -1 {
				node.Left = &TreeNode{
					Val: nums[ind],
				}
				tempNodes = append(tempNodes, node.Left)
			}
			ind++
			if ind >= length {
				break
			}
			if nums[ind] != -1 {
				node.Right = &TreeNode{
					Val: nums[ind],
				}
				tempNodes = append(tempNodes, node.Right)
			}
			ind++
		}
		nodeLevel = tempNodes
	}
	return root
}

func GetNodesByPre(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	curNodeList := make([]*TreeNode, 1)
	curNodeList[0] = root
	curNodeList = append(curNodeList, GetNodesByPre(root.Left)...)
	curNodeList = append(curNodeList, GetNodesByPre(root.Right)...)
	return curNodeList
}

type nodeWithParent struct {
	attr   bool //属性，false-left，true-right
	cur    *TreeNode
	left   *nodeWithParent
	right  *nodeWithParent
	parent *nodeWithParent
}

func ConstructNewTreeWithParent(root *TreeNode, parentNode *nodeWithParent, attr bool) *nodeWithParent {
	if root == nil {
		return nil
	}
	res := &nodeWithParent{
		parent: parentNode,
		cur:    root,
		attr:   attr,
	}
	res.left = ConstructNewTreeWithParent(root.Left, res, false)
	res.right = ConstructNewTreeWithParent(root.Right, res, true)
	return res
}

func GetLeafNodes(root *nodeWithParent) []*nodeWithParent {
	if root == nil {
		return nil
	}
	if root.left == nil && root.right == nil {
		return []*nodeWithParent{root}
	}
	res := make([]*nodeWithParent, 0)
	res = append(res, GetLeafNodes(root.left)...)
	res = append(res, GetLeafNodes(root.right)...)
	return res
}

/*问题*/
/*
给定一个数组，[10,5,15,3,7,-1,18]，属于二叉树的层次遍历，其中-1表示该节点为null
重构二叉树
*/
/*思路*/
/*

 */
func ConstructTreeInLevel(nums []int) (root *TreeNode) {
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
			var left *TreeNode = nil
			if nums[2*i-1] != -1 {
				left = &TreeNode{
					Val: nums[2*i-1],
				}
			}
			if curPos != nil {
				curPos.Left = left
			}
			tempList.PushBack(left)
		}
		if (i * 2) < length {
			var right *TreeNode = nil
			if nums[2*i] != -1 {
				right = &TreeNode{
					Val: nums[2*i],
				}
			}
			if curPos != nil {
				curPos.Right = right
			}
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

/*问题*/
/*
给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。

二叉搜索树保证具有唯一的值。



示例 1：

输入：root = [10,5,15,3,7,null,18], L = 7, R = 15
输出：32
示例 2：

输入：root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
输出：23


提示：

树中的结点数量最多为 10000 个。
最终的答案保证小于 2^31。
*/
/*思路*/
/*
最直接的思路，遍历二叉树，只要是>=7 && <=15的节点一定在L和R之间
如果某节点<7,则该节点的左节点都不用遍历了
如果某节点>15,则该节点的右节点都不用遍历了
*/
func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	if root == nil {
		return res
	}

	if root.Val >= L && root.Val <= R {
		res = res + root.Val
		temp1 := rangeSumBST(root.Left, L, R)
		temp2 := rangeSumBST(root.Right, L, R)
		res = res + temp1 + temp2
	} else if root.Val < L {
		temp2 := rangeSumBST(root.Right, L, R)
		res = res + temp2
	} else if root.Val > R {
		temp1 := rangeSumBST(root.Left, L, R)
		res = res + temp1
	}
	return res
}

/*问题*/
/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
*/
/*思路*/
/*
1.从任意节点来看，其左右节点都经过了翻转
2.从上至下，要翻转当前节点，先翻转当前节点的左右子节点
3.构建递归，递归终止条件，某节点为空，则不需要翻转，直接返回
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	return root
}

/*问题*/
/*
给定二叉树的根节点 root，找出存在于不同节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。

（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）



示例：
			8
           / \
          3   10
         / \    \
        1   6    14
           / \   /
          4   7 13

输入：[8,3,10,1,6,null,14,null,null,4,7,13]
      0,1,2  3,4  5,   6, 7 ,  8,  9 ,10,11
输出：7
解释：
我们有大量的节点与其祖先的差值，其中一些如下：
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。


提示：

树中的节点数在 2 到 5000 之间。
每个节点的值介于 0 到 100000 之间。
*/
/*思路*/
/*
1.之前没考虑到给出的是树结构的数据
输入的形式还是完全二叉树的形式，可以根据公式
父 = floor((子-1)/2)
从第1位开始计算，只算该子节点与以上祖先节点的最大值

2.后序遍历，需要一个数组存储当前已经走到，但子节点还没走完的祖先数组
思维僵化：
其实可以在遍历过程仅保存上下界，然后跟当前节点比较，而不是每次都跟所有祖先比较
*/
func postSearchTreeWithAncestor(root *TreeNode, ancestor []*TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiff := 0
	if len(ancestor) != 0 {
		for _, node := range ancestor {
			tempDiff := node.Val - root.Val
			if tempDiff < 0 {
				tempDiff = tempDiff * -1
			}
			if tempDiff > maxDiff {
				maxDiff = tempDiff
			}
		}
	}
	ancestor = append(ancestor, root)
	leftMaxDiff := postSearchTreeWithAncestor(root.Left, ancestor)
	rightMaxDiff := postSearchTreeWithAncestor(root.Right, ancestor)
	ancestor = ancestor[:len(ancestor)-1]

	if maxDiff < leftMaxDiff {
		maxDiff = leftMaxDiff
	}
	if maxDiff < rightMaxDiff {
		maxDiff = rightMaxDiff
	}

	return maxDiff
}

func MaxAncestorDiff(root *TreeNode) int {
	ancestor := make([]*TreeNode, 0)
	return postSearchTreeWithAncestor(root, ancestor)
}

/*问题*/
/*
给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。

例如，从根到叶子节点路径 1->2->3 代表数字 123。

计算从根到叶子节点生成的所有数字之和。

说明: 叶子节点是指没有子节点的节点。

示例 1:

输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.
示例 2:

输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-root-to-leaf-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
还是遍历就完事了，只是需要存一下当前节点的祖先已经构成的数据的大小
遍历方式，先序遍历就行了
*/
func SumNumbersWithAncestors(root *TreeNode, ancestors *int, curRes *int) {
	if root == nil {
		return
	}
	*ancestors = *ancestors*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*curRes = *curRes + *ancestors
	}
	SumNumbersWithAncestors(root.Left, ancestors, curRes)
	SumNumbersWithAncestors(root.Right, ancestors, curRes)
	*ancestors = *ancestors / 10
	return
}

func SumNumbers(root *TreeNode) int {
	ancestors := 0
	curRes := 0
	SumNumbersWithAncestors(root, &ancestors, &curRes)
	return curRes
}

/*问题*/
/*
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
换种说法，思路比较清晰
遍历所有节点，计算从某节点出发，和最大的左子树和 与 右子树和 以及 两子树和相加(均包括该节点本身吧)
*/
func getTwoPartSum(root *TreeNode, curMax *int, isSet *bool) (leftSum int, rightSum int) {
	if root == nil {
		return 0, 0
	}
	ll, lr := getTwoPartSum(root.Left, curMax, isSet)
	if ll < lr {
		ll = lr
	}
	rl, rr := getTwoPartSum(root.Right, curMax, isSet)
	if rl < rr {
		rl = rr
	}
	if ll > 0 {
		leftSum = ll + root.Val
	} else {
		leftSum = root.Val
	}
	if rl > 0 {
		rightSum = rl + root.Val
	} else {
		rightSum = root.Val
	}

	tempMax := int(math.Max(float64(leftSum+rightSum-root.Val), math.Max(float64(leftSum), float64(rightSum))))
	tempMax = int(math.Max(float64(tempMax), float64(root.Val)))
	if !*isSet {
		*curMax = tempMax
		*isSet = true
	} else {
		if *curMax < tempMax {
			*curMax = tempMax
		}
	}
	return leftSum, rightSum
}
func MaxPathSum(root *TreeNode) int {
	max := new(int)
	isSet := new(bool)
	*isSet = false
	getTwoPartSum(root, max, isSet)
	return *max
}

/*问题*/
/*
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

两棵树重复是指它们具有相同的结构以及相同的结点值。

示例 1：

        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4
下面是两个重复的子树：

      2
     /
    4
和

    4
因此，你需要以列表的形式返回上述重复子树的根结点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-duplicate-subtrees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
子树重复，首先得保证叶子节点重复
需要一个数组存储叶子结点
对叶子节点的数值进行分类，判断是否相等
然后进行反向遍历
*/
func GetLevelLoopString(nodes []*TreeNode) string {
	res := ""
	if len(nodes) == 0 {
		return res
	}
	nextNodes := make([]*TreeNode, 0)
	for _, node := range nodes {
		if node == nil {
			res = res + "n" + "."
			continue
		}
		res = res + strconv.Itoa(node.Val) + "."
		nextNodes = append(nextNodes, node.Left)
		nextNodes = append(nextNodes, node.Right)
	}
	res = res + GetLevelLoopString(nextNodes)
	return res
}

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type tempStruct struct {
		node        *TreeNode
		checkRepeat bool
	}

	checkRes := make(map[string]tempStruct)

	nodeList := GetNodesByPre(root)
	for _, node := range nodeList {
		tempStr := GetLevelLoopString([]*TreeNode{node})
		if _, ok := checkRes[tempStr]; ok {
			checkRes[tempStr] = tempStruct{
				node:        node,
				checkRepeat: true,
			}
		} else {
			checkRes[tempStr] = tempStruct{
				node:        node,
				checkRepeat: false,
			}
		}
	}

	res := make([]*TreeNode, 0)
	for _, checkNode := range checkRes {
		if checkNode.checkRepeat {
			res = append(res, checkNode.node)
		}
	}
	return res
}
