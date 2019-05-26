package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	//nums := []int{2, 4, 7, 3, 1, 6, 8, 33, 11, 2, 5, 7}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	SelectSort(nums)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	//nums := []int{2, 4, 7, 3, 1, 6, 8, 33, 11, 2, 5, 7}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	fmt.Println(BubbleSort(nums))
}

func TestQuickSort(t *testing.T) {
	//nums := []int{2, 4, 7, 3, 1, 6, 8, 33, 11, 2, 5, 7}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	//nums := []int{4, 2, 7, 3, 1, 6}
	QuickSort(nums)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	//nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	nums := []int{4, 2, 7, 3, 1, 6, 11}
	//nums := []int{2, 1}
	res := MergeSort(nums)
	fmt.Println(nums, res)
}
