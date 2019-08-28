package reservoir_sampling

import (
	"fmt"
	"github.com/bluesky1024/goAlgorithm/linked_list"
	"testing"
)

func TestSolution_GetRandom(t *testing.T) {
	a := &linked_list.ListNode{
		Val: 1,
		Next: &linked_list.ListNode{
			Val: 2,
			Next: &linked_list.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	s := SolutionConstructor(a)

	cnt := make([]int, 3)
	for i := 1; i < 100000; i++ {
		cnt[s.GetRandom()-1]++
	}
	fmt.Println(cnt)
}
