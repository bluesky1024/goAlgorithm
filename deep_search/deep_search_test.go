package deep_search

import (
	"fmt"
	"github.com/bluesky1024/goAlgorithm/tree"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	rooms := [][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}
	res := CanVisitAllRooms(rooms)
	fmt.Println(res)
}

func TestFindRotateSteps(t *testing.T) {

}

func TestBinaryTreePaths(t *testing.T) {
	root := tree.ConstrucTreeInLevelWithoutInvalidNode([]int{1, 2, 3, -1, 5})
	fmt.Println(BinaryTreePaths(root))
}
