package dynamic

import (
	"fmt"
	"testing"
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
	trangles := [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}
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
