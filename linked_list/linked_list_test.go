package linked_list

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

}

func TestSortList(t *testing.T) {
	head := &ListNode{
		Val:  4,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	pCur := head
	for {
		if pCur == nil {
			break
		}
		fmt.Println(pCur.Val)
		pCur = pCur.Next
	}
	fmt.Println("after")
	SortList(head)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}
