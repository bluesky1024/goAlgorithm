package linked_list

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
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
