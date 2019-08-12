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

func TestNumSquaresV1(t *testing.T) {
	fmt.Println(NumSquaresV1(4703))
}

func TestNumSquaresV2(t *testing.T) {
	fmt.Println(NumSquaresV2(5))
}
