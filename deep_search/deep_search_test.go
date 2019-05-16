package deep_search

import (
	"fmt"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	rooms := [][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}
	res := CanVisitAllRooms(rooms)
	fmt.Println(res)
}
