package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestMergeTwoList(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  11,
							Next: nil,
						},
					},
				},
			},
		},
	}

	b := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 10,
						Next: &ListNode{
							Val:  12,
							Next: nil,
						},
					},
				},
			},
		},
	}

	c := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 12,
					Next: &ListNode{
						Val: 15,
						Next: &ListNode{
							Val:  18,
							Next: nil,
						},
					},
				},
			},
		},
	}

	d := MergeKLists([]*ListNode{a, b, c})
	d.Prinf()
}

func TestLink(t *testing.T) {
	mLink := NewLinkNode([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, mLink.PrintAndOutput(), []int{1, 2, 3, 4, 5, 6})
	mLink.Del(1)
	assert.Equal(t, mLink.PrintAndOutput(), []int{2, 3, 4, 5, 6})
	mLink.Del(4)
	assert.Equal(t, mLink.PrintAndOutput(), []int{2, 3, 5, 6})

	mLink.Insert(0, 10)
	assert.Equal(t, mLink.PrintAndOutput(), []int{10, 2, 3, 5, 6})
	mLink.Insert(2, 20)
	assert.Equal(t, mLink.PrintAndOutput(), []int{10, 2, 20, 3, 5, 6})

	mLink.Reverse()
	assert.Equal(t, mLink.PrintAndOutput(), []int{6, 5, 3, 20, 2, 10})
}

func TestLink_ReverseEveryN(t *testing.T) {
	mLink := NewLinkNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
	mLink.ReverseEveryN(1)
	assert.Equal(t, mLink.PrintAndOutput(), []int{1, 2, 3, 4, 5, 6, 7, 8})

	mLink = NewLinkNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
	mLink.ReverseEveryN(2)
	assert.Equal(t, mLink.PrintAndOutput(), []int{2, 1, 4, 3, 6, 5, 8, 7})

	mLink = NewLinkNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
	mLink.ReverseEveryN(3)
	assert.Equal(t, mLink.PrintAndOutput(), []int{3, 2, 1, 6, 5, 4, 8, 7})

	mLink = NewLinkNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
	mLink.ReverseEveryN(4)
	assert.Equal(t, mLink.PrintAndOutput(), []int{4, 3, 2, 1, 8, 7, 6, 5})

	mLink = NewLinkNode([]int{1, 2, 3})
	mLink.ReverseEveryN(4)
	assert.Equal(t, mLink.PrintAndOutput(), []int{3, 2, 1})
}

func Test_partition(t *testing.T) {
	l := NewListNode([]int{1, 4, 3, 0, 2, 5, 2})
	l.Prinf()

	n := partition(l, 3)
	n.Prinf()

	l1 := NewListNode([]int{1, 4, 3, 2, 5, 2})
	l1.Prinf()

	n1 := partition(l1, 3)
	n1.Prinf()

	l2 := NewListNode([]int{4, 3, 2, 1, 5, 2})
	l2.Prinf()

	n2 := partition(l2, 3)
	n2.Prinf()

	l3 := NewListNode([]int{4, 3, 5})
	l3.Prinf()

	n3 := partition(l3, 3)
	n3.Prinf()
}
