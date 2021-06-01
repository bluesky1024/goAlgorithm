package bit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	fmt.Println(HammingDistance(1, 4))
}

func TestFindComplement(t *testing.T) {
	fmt.Println(FindComplement(5))
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	fmt.Println(FindRepeatedDnaSequences("AAAAAAAAAAA"))
}

func TestHammingDistance2(t *testing.T) {
	var nums uint32 = 11
	fmt.Println(HammingWeight(nums))
}

func Test_swapNumbers(t *testing.T) {
	assert.Equal(t, swapNumbers([]int{1, 2}), []int{2, 1})
}

func Test_getSum(t *testing.T) {
	assert.Equal(t, getSum(1, 2), 3)
	assert.Equal(t, getSum(-2, 3), 1)
}
