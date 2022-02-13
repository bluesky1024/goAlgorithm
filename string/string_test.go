package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(UniqueMorseRepresentations(words))
}

func TestRestoreIpAddresses(t *testing.T) {
	str := "25525511135"
	fmt.Println(RestoreIpAddresses(str))
}

func TestAdd(t *testing.T) {
	num1 := "660"
	num2 := "66"
	fmt.Println(Add(num1, num2))
}

func TestMultiply(t *testing.T) {
	num1 := "22"
	num2 := "33"
	fmt.Println(Multiply(num1, num2))
}

func TestAmbiguousCoordinates(t *testing.T) {
	fmt.Println(AmbiguousCoordinates("(123)"))
}

func TestConvert(t *testing.T) {
	fmt.Println(Convert("A", 1))
}

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	ReverseString(s)
	fmt.Printf("%c", s)
}

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	fmt.Println(s[3:])
	t.Log(ReverseWords(s))
}

func TestCheckPalindrome(t *testing.T) {
	s := "a"
	fmt.Println(CheckPalindrome([]byte(s)))
}

func TestCountSubstrings(t *testing.T) {
	s := "aaa"
	fmt.Println(CountSubstrings_v2(s))
}

func TestScoreOfParentheses(t *testing.T) {
	s := "(()(()))(()())"
	//s = "(((())))"
	fmt.Println(ScoreOfParentheses(s))
}

func TestDefangIPaddr(t *testing.T) {
	fmt.Println(DefangIPaddr("1.1.1.1"))
}

func TestLetterCombinations(t *testing.T) {
	nums := "2345"
	fmt.Println(LetterCombinations(nums))
}

func TestOptimalDivision(t *testing.T) {
	assert.Equal(t, OptimalDivision([]int{1000, 100, 10, 2}), "1000/(100/10/2)")
}

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("babad"), "bab")
	assert.Equal(t, longestPalindrome("cbbd"), "bb")
	assert.Equal(t, longestPalindrome("a"), "a")
	assert.Equal(t, longestPalindrome("aaaa"), "aaaa")
}

func Test_generateParenthesis(t *testing.T) {
	assert.Len(t, generateParenthesis(1), 1)
	assert.Len(t, generateParenthesis(2), 2)
	assert.Len(t, generateParenthesis(3), 5)
}

func Test_countAndSay(t *testing.T) {
	assert.Equal(t, countAndSaySub("111221"), "312211")

	assert.Equal(t, countAndSay(1), "1")
	assert.Equal(t, countAndSay(2), "11")
	assert.Equal(t, countAndSay(3), "21")
	assert.Equal(t, countAndSay(4), "1211")
	assert.Equal(t, countAndSay(5), "111221")
	assert.Equal(t, countAndSay(6), "312211")
}

func Test_simplifyPath(t *testing.T) {
	path := "/a/./b/../../c"
	t.Log(simplifyPath(path))
}

func Test_ReverseNum(t *testing.T) {
	assert.Equal(t, ReverseNum(1516000), "0006151")
	assert.Equal(t, ReverseNum(0), "0")
}

func Test_isScramble(t *testing.T) {
	// assert.True(t, isScramble("great", "rgeat"))
	// assert.False(t, isScramble("abcde", "caebd"))
	// assert.True(t, isScramble("a", "a"))

	assert.True(t, isScramble("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"))
}
