package hash

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	nums := []int{0, 0}
	fmt.Println(HIndex(nums))
}

func TestIslandPerimeter(t *testing.T) {
	grid := [][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}}
	fmt.Println(IslandPerimeter(grid))
}

func TestLargestValsFromLabels(t *testing.T) {
	v := []int{3, 0, 3, 0, 6}
	l := []int{0, 2, 1, 1, 0}
	fmt.Println(LargestValsFromLabels(v, l, 4, 1))
}

func TestFourSum(t *testing.T) {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	fmt.Println(FourSum(nums, 0))
}
