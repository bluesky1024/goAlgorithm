package dynamic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	myNumArray := ConstructorNumArray(nums)
	fmt.Println(myNumArray.SumRange(2, 5))
}

func TestMinPathSum(t *testing.T) {
	minSumArr := make([][]int, 3)
	temp1 := []int{1, 3, 1}
	temp2 := []int{1, 5, 1}
	temp3 := []int{4, 2, 1}
	minSumArr[0] = temp1
	minSumArr[1] = temp2
	minSumArr[2] = temp3
	fmt.Println(MinPathSum(minSumArr))
}

func TestNthUglyNumber(t *testing.T) {
	fmt.Println(NthUglyNumber(10))
}

func TestMinimumTotal(t *testing.T) {
	trangles := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(MinimumTotal(trangles))
}

func TestMinCut(t *testing.T) {
	s := "apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"
	fmt.Println(MinCut(s))
}

func TestWordBreak(t *testing.T) {
	strs := []string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"}
	fmt.Println(WordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", strs))

	//strs := []string{"go", "php", "c++", "java", "ba"}
	//fmt.Println(WordBreak("java1go", strs))

}

func TestTallestBillboard(t *testing.T) {
	//rods := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 50, 50, 50, 150, 150, 150, 100, 100, 100, 123}
	//fmt.Println(TallestBillboard(rods))

	//rods := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(TallestBillboard(rods))

	rods := []int{94, 98, 87, 111, 86, 97, 101, 110, 100, 91, 98, 106, 114, 109, 107, 87, 104}
	fmt.Println(TallestBillboard(rods))
}

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	fmt.Println(CountNumbersWithUniqueDigits(13))
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(MaxProfit([]int{7, 6, 5, 4, 3, 2, 1}))

	fmt.Println(MaxProfit([]int{2, 1, 4}))
}

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)
	assert.Equal(t, LargestRectangleArea([]int{2, 1, 2}), 3)

	assert.Equal(t, LargestRectangleAreaFormal([]int{2, 1, 5, 6, 2, 3}), 10)
	assert.Equal(t, LargestRectangleAreaFormal([]int{2, 1, 2}), 3)
}

func Test_generate(t *testing.T) {
	assert.Equal(t, generate(1), [][]int{{1}})
	assert.Equal(t, generate(3), [][]int{{1}, {1, 1}, {1, 2, 1}})
}

func Test_rob(t *testing.T) {
	assert.Equal(t, rob([]int{1, 2, 3, 1}), 4)
	assert.Equal(t, rob([]int{2, 7, 9, 3, 1}), 12)
	assert.Equal(t, rob([]int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124}), 7102)
}

func Test_superUgly(t *testing.T) {
	assert.Equal(t, superUgly(12, []int{2, 7, 13, 19}), 32)
}

func Test_findMaxForm(t *testing.T) {
	//assert.Equal(t, findMaxForm([]string{"10", "0", "1"}, 1, 1), 2)
	assert.Equal(t, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3), 4)
}
