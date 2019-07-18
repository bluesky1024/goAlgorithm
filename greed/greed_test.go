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

func TestTwoCitySchedCost(t *testing.T) {
	costs := [][]int{[]int{259, 770}, []int{448, 54}, []int{926, 667}, []int{184, 139}, []int{840, 118}, []int{577, 469}}
	fmt.Println(TwoCitySchedCost(costs))
}

func TestMinAddToMakeValid(t *testing.T) {
	s := "(())())"
	fmt.Println(MinAddToMakeValid(s))
}
