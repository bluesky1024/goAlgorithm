package design

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	myStack := MyStackConstructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
}
