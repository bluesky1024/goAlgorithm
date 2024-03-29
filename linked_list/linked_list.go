package linked_list

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Prinf() {
	cur := l
	for cur != nil {
		fmt.Printf("%p\n", cur)
		cur = cur.Next
	}
	cur = l
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("---")
}

/*问题*/
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
/*思路*/
/*
for循环从前往后遍历，有就相加，没有就取其中一个值
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	p1 := l1
	p2 := l2
	temp1 := 0
	temp2 := 0
	isAdd := false

	var res *ListNode
	var pRes *ListNode

	for {
		if p1 == nil && p2 == nil {
			break
		}
		temp1 = 0
		temp2 = 0
		if p1 != nil {
			temp1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			temp2 = p2.Val
			p2 = p2.Next
		}
		val := temp1 + temp2
		if isAdd {
			val++
			isAdd = false
		}
		if val >= 10 {
			isAdd = true
		}

		listTemp := &ListNode{
			Val:  val % 10,
			Next: nil,
		}

		if pRes == nil {
			res = listTemp
			pRes = res
		} else {
			pRes.Next = listTemp
			pRes = pRes.Next
		}
	}
	if isAdd {
		pRes.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return res
}

/*问题*/
/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
调用原生的sort.Sort()
*/
type NewList struct {
	Sub *ListNode
}

func (l *NewList) Swap(i, j int) {
	length := 0
	pCur := l.Sub
	var pI *ListNode
	var pJ *ListNode
	for {
		if pCur == nil {
			break
		}
		if length == i {
			pI = pCur
		}
		if length == j {
			pJ = pCur
		}
		length++
		pCur = pCur.Next
		if pI != nil && pJ != nil {
			break
		}
	}
	pI.Val, pJ.Val = pJ.Val, pI.Val
	return
}

func (l *NewList) Len() int {
	length := 0
	pCur := l.Sub
	for {
		if pCur == nil {
			break
		}
		length++
		pCur = pCur.Next
	}
	return length
}

func (l *NewList) Less(i, j int) bool {
	length := 0
	pCur := l.Sub
	var pI *ListNode
	var pJ *ListNode
	for {
		if pCur == nil {
			break
		}
		if length == i {
			pI = pCur
		}
		if length == j {
			pJ = pCur
		}
		length++
		pCur = pCur.Next
		if pI != nil && pJ != nil {
			break
		}
	}
	if pI.Val < pJ.Val {
		return true
	}
	return false
}

func SortList(head *ListNode) *ListNode {
	tempNewList := &NewList{
		Sub: head,
	}
	sort.Sort(tempNewList)
	head = tempNewList.Sub
	return head
}

/*问题*/
/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
交换流程：
输出第二个节点
保存原第二个节点的next到Next变量
第二个节点的next指向第一个节点
第一个节点的next指向上述Next节点
*/
func swapTwoNodes(first *ListNode) (newFirst *ListNode, next *ListNode) {
	if first == nil || first.Next == nil {
		return first, nil
	}
	newFirst = first.Next
	next = newFirst.Next
	newFirst.Next = first
	newFirst.Next.Next = next
	return newFirst, next
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res, next := swapTwoNodes(head)
	cur := res

	for {
		if next == nil {
			break
		}
		var temp *ListNode
		temp, next = swapTwoNodes(next)
		cur.Next.Next = temp
		cur = temp
	}

	return res
}

/*问题*/
/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/odd-even-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
重点在于
1.删除节点
a,b,c之间删除b
a.Next = c
2.插入节点
a,b中插入c
a.Next = c
c.Next = b
*/
func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	p1 := head

	for {
		//奇数前挪，偶数不动
		if p == nil {
			break
		}

		//第一个是偶数
		p = p.Next

		//第二个是奇数
		if p == nil || p.Next == nil {
			break
		}
		//要把如下temp拆出来，挪到p1后面
		temp := p.Next
		p.Next = temp.Next

		temp2 := p1.Next
		p1.Next = temp
		temp.Next = temp2

		p1 = temp
	}
	return head
}

/*问题*/
/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1.迭代，正向遍历，顺便逐个反转
2.递归，求出head.Next的反转结果，然后与head拼接
*/
func ReverseListV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head
	cur := head.Next
	res.Next = nil
	for {
		if cur == nil {
			break
		}
		temp1 := res
		temp2 := cur.Next
		res = cur
		res.Next = temp1
		cur = temp2
	}
	return res
}

func ReverseListV2(head *ListNode) (res *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	res = ReverseListV2(head.Next)
	cur := res
	for {
		if cur.Next == nil {
			cur.Next = head
			cur.Next.Next = nil
			break
		}
		cur = cur.Next
	}

	return res
}

/*问题*/
/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
*/
/*思路*/
/*
简化问题为两个链表的合并
然后建立循环，把第2个到最后一个合并到第一个链表里
*/
func MergeTwoList(la *ListNode, lb *ListNode) *ListNode {
	if la == nil {
		return lb
	}

	if lb == nil {
		return la
	}

	pa := la
	pb := lb
	var res *ListNode
	var pLast *ListNode
	if pa.Val <= pb.Val {
		res = pa
		pLast = pa
		pa = pa.Next
	} else {
		res = pb
		pLast = pb
		pb = pb.Next
	}
	for {
		if pa != nil && pb != nil {
			var temp *ListNode
			if pa.Val <= pb.Val {
				temp = pa
				pa = pa.Next
			} else {
				temp = pb
				pb = pb.Next
			}

			pLast.Next = temp
			pLast = temp
			continue
		}

		if pa != nil && pb == nil {
			pLast.Next = pa
			pLast = pa
			pa = pa.Next
			continue
		}

		if pa == nil && pb != nil {
			pLast.Next = pb
			pLast = pb
			pb = pb.Next
			continue
		}

		break
	}

	return res
}
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	l1 := lists[0]
	for i := 1; i < len(lists); i++ {
		temp := MergeTwoList(l1, lists[i])
		l1 = temp
	}

	return l1
}

// Link 链表实现
type Link struct {
	root *linkNode
}

// linkNode 链表节点
type linkNode struct {
	data int
	next *linkNode
}

func NewLinkNode(nums []int) *Link {
	if len(nums) == 0 {
		return nil
	}
	root := &linkNode{
		data: nums[0],
		next: nil,
	}
	cur := root
	for i := 1; i < len(nums); i++ {
		cur.next = &linkNode{
			data: nums[i],
			next: nil,
		}
		cur = cur.next
	}
	return &Link{root: root}
}

// Print 按顺序打印元素值并输出
func (l *Link) PrintAndOutput() []int {
	res := make([]int, 0)
	cur := l.root
	for {
		if cur == nil {
			break
		}
		res = append(res, cur.data)
		cur = cur.next
	}
	fmt.Println(res)
	return res
}

// Del 删除特定数据
func (l *Link) Del(t int) {
	cur := l.root
	var pre *linkNode
	for {
		if cur == nil {
			break
		}
		if cur.data == t {
			if pre == nil {
				l.root = cur.next
				return
			}
			pre.next = cur.next
			return
		}
		pre = cur
		cur = cur.next
	}
}

// Insert 在指定pos插入数据t，pos从0开始
func (l *Link) Insert(pos int, t int) {
	curPos := 0
	cur := l.root
	var pre *linkNode
	for {
		if cur == nil {
			break
		}
		if curPos == pos {
			if pre == nil {
				l.root = &linkNode{
					data: t,
					next: cur,
				}
				return
			}
			pre.next = &linkNode{
				data: t,
				next: cur,
			}
			return
		}
		pre = cur
		cur = cur.next
		curPos++
	}
}

// Reverse 反转
func (l *Link) Reverse() {
	cur := l.root
	var pre *linkNode
	for {
		if cur == nil {
			break
		}
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}
	l.root = pre
}

// ReverseEveryThree 每隔n个反转一次
// 例：n=3,[1,2,3,4,5,6,7,8] => [3,2,1,6,5,4,8,7]
// 首先构造正常的全量反转，并引入遍历计数
// 当计数为3时 cur-3 pre-2 temp-4
func (l *Link) ReverseEveryN(n int) {
	if n <= 1 {
		return
	}
	curCnt := 0
	cur := l.root
	l.root = nil
	var pre *linkNode

	lastEndList := []*linkNode{}
	lastEnd := cur
	for {
		curCnt++
		if cur == nil {
			break
		}
		temp := cur.next
		cur.next = pre
		pre = cur

		if curCnt == 1 {
			lastEndList = append(lastEndList, cur)
			lastEnd = cur
		}

		if curCnt == n {
			if l.root == nil {
				l.root = cur
			} else {
				lastEnd = lastEndList[0]
				lastEndList = lastEndList[1:]
				lastEnd.next = cur
			}
			pre = nil
			curCnt = 0
		}

		cur = temp
	}

	if l.root != nil {
		lastEndList[0].next = pre
	} else {
		l.root = pre
	}
}

/*问题*/
/*

给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。



示例 1：


输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]
示例 2：

输入：head = [2,1], x = 2
输出：[1,2]


提示：

链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
/*思路*/
/*
从前往后找到第一个大于等于x的数，保存其ptr
继续往后遍历，发现小于x的数就插到ptr前面，同时更新这个ptr的值为当前这个数的ptr
后续再发现小于x的数，就插入ptr后面
就是写的太丑陋了。。。
*/

func NewListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val: nums[0],
	}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{
			Val: nums[i],
		}
		cur = cur.Next
	}
	return head
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	mvBehind := func(pos *ListNode, oriBefore *ListNode, t *ListNode) *ListNode {
		oriBefore.Next = t.Next
		if pos == nil {
			return t
		}
		t.Next = pos.Next
		pos.Next = t
		return nil
	}

	oriHead := head
	var last *ListNode
	var insertPoint *ListNode
	findFirstOver := false
	for {
		if head == nil {
			break
		}
		if head.Val >= x {
			if !findFirstOver {
				findFirstOver = true
				insertPoint = last
			}

			last = head
			head = head.Next
		} else {
			if !findFirstOver {
				last = head
				head = head.Next
			} else {
				tmp := head.Next
				node := mvBehind(insertPoint, last, head)
				insertPoint = head
				if node != nil {
					node.Next = oriHead
					oriHead = node
				}
				head = tmp
			}
		}
	}

	return oriHead
}
