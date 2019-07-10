package greed

import (
	"fmt"
	"testing"
)

func TestMinDeletionSize(t *testing.T) {
	a := []string{"cba", "daf", "ghi"}
	fmt.Println(MinDeletionSize(a))
}

func TestLastStoneWeight(t *testing.T) {
	res := LastStoneWeight([]int{3, 2, 6, 4})
	fmt.Println(res)
}
