package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinAjust(t *testing.T) {
	nums := []int{5, 1, 13, 3, 16, 7, 10, 14, 6, 9}
	MinAjust(nums)
	fmt.Println(nums)
}

func TestKthLargest_Add(t *testing.T) {
	k := 3
	nums := []int{14, 15, 8, 12, 6}
	kthLargest := ConstructorKthLargest(k, nums)
	fmt.Println(kthLargest.Add(13))
	fmt.Println(kthLargest.Add(19))
	fmt.Println(kthLargest.Add(22))
}

func Test_checkBrackets(t *testing.T) {
	assert.Equal(t, checkBrackets("{{{}()}}"), true)
	assert.Equal(t, checkBrackets("{([]{})}"), false)
	assert.Equal(t, checkBrackets("{([]{})}"), false)
	assert.Equal(t, checkBrackets("{}{}[()]"), true)
	assert.Equal(t, checkBrackets("{([)]} "), false)
}
