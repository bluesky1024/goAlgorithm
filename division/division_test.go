package division

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {

}

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(Search(nums, 2))
}

func TestSearchV2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12, 22, 33, 41, 63, 84, 111, 222, 333, 555, 3331}
	fmt.Println(SearchV2(nums, 85))
}

func TestFindMin(t *testing.T) {
	nums := []int{7, 0, 1, 2, 3, 5, 6}
	fmt.Println(FindMin(nums))
}

func TestFindMinV2(t *testing.T) {
	nums := []int{10, 1, 10, 10, 10}
	fmt.Println(FindMinV2(nums))
}
