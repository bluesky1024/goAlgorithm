package linked_list

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
