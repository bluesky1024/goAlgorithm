package string

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
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
	fmt.Println(s[3:len(s)])
	return
	fmt.Println(ReverseWords(s))
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
