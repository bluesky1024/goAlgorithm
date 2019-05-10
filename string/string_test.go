package string

import (
	"fmt"
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
