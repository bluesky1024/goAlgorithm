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

func TestFindMin(t *testing.T) {
	nums := []int{7, 0, 1, 2, 3, 5, 6}
	fmt.Println(FindMin(nums))
}
