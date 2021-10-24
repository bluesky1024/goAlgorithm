package breadth_search

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestXOChessSolve(t *testing.T) {
	board := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	XOChessSolve(board)
	fmt.Println(board)
}

func TestMakeConnected(t *testing.T) {
	//assert.Equal(t, MakeConnectedV2(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}), 2)
	//assert.Equal(t, MakeConnectedV2(5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}), 0)
	assert.Equal(t, MakeConnected(12, [][]int{{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2}}), 4)
}

func Test_exist(t *testing.T) {
	assert.True(t, exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}, "ABCCED"))

	//assert.False(t, exist([][]byte{
	//	[]byte("ABCE"),
	//	[]byte("SFES"),
	//	[]byte("ADEE"),
	//}, "ABCB"))
}
