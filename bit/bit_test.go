package bit

import (
	"fmt"
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
