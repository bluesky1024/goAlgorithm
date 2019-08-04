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
