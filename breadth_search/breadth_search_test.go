package breadth_search

import (
	"fmt"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	inputNums := [][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}
	res := OrangesRotting(inputNums)
	fmt.Println(res)
}
