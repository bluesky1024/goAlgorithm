package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, nums, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestMergeSort(t *testing.T) {
	//nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	nums := []int{4, 2, 7, 3, 1, 6, 11}
	//nums := []int{2, 1}
	res := MergeSort(nums)
	fmt.Println(nums, res)
}

func TestHeapSort(t *testing.T) {
	nums := []int{35, 14, 75, 51, 133, 21, 3, 12}
	fmt.Println(nums)
	HeapSort(nums)
	fmt.Println(nums)
}

func TestLargestPerimeter(t *testing.T) {
	fmt.Println(LargestPerimeter([]int{3, 6, 2, 3}))
}

func TestPancakeSort(t *testing.T) {
	fmt.Println(PancakeSort([]int{3, 2, 4, 1}))
}

func TestGetStrongest(t *testing.T) {
	assert.Equal(t, GetStrongest([]int{6, 7, 11, 7, 6, 8}, 5), []int{11, 8, 6, 6, 7})
}
