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

func TestSwapPairs(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	a.Prinf()
	a = SwapPairs(a)
	fmt.Println("after swap")
	a.Prinf()
}

func TestOddEvenList(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	a.Prinf()
	a = OddEvenList(a)
	fmt.Println("after odd")
	a.Prinf()
}

func TestReverseListV1(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	a.Prinf()
	a = ReverseListV1(a)
	fmt.Println("after reverse v1")
	a.Prinf()

	a = ReverseListV2(a)
	fmt.Println("after reverse v2")
	a.Prinf()
}
