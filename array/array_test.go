package array

import (
	"fmt"
	"testing"
)

func TestFindNearestNum(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(FindNearestNum(nums, 0, len(nums)-1, 6))
}

func TestFlipAndInvertImage(t *testing.T) {
	nums := [][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}
	nums = FlipAndInvertImage(nums)
	fmt.Println(nums)
}

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(SortedSquares(nums))
}

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(ProductExceptSelf(nums))
}

func TestFindInMountainArray(t *testing.T) {
	nums := MountainArray{
		Arr: []int{0, 5, 3, 1},
	}
	fmt.Println(FindMoutain(&nums, 0, nums.length()-1, nums.get(0), nums.get(nums.length()-1)))
	fmt.Println(FindInMountainArray(1, &nums))
}

func TestMaxPoints(t *testing.T) {
	points := [][]int{[]int{1, 1}, []int{1, 1}, []int{2, 3}}
	//points := [][]int{[]int{1, 1}, []int{3, 2}, []int{5, 3}, []int{4, 1}, []int{2, 3}, []int{1, 4}}
	fmt.Println(MaxPoints(points))
}

func TestIsUgly(t *testing.T) {
	fmt.Println(IsUgly(7))
}

func TestGameOfLifeV1(t *testing.T) {
	nums := [][]int{[]int{0, 1, 0, 0, 1, 1, 0}, []int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 0, 0, 0, 0, 1}, []int{1, 1, 0, 0, 1, 0, 0}}
	GameOfLifeV1(nums)
	fmt.Println(nums)
}
