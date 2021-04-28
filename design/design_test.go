package design

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMyStack(t *testing.T) {
	myStack := MyStackConstructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
}

func TestMyLinkedList(t *testing.T) {
	//fmt.Println("init")
	//linkedList := ConstructorMyLinkedList()
	//linkedList.Print()
	//fmt.Println("add at head 1")
	//linkedList.AddAtHead(1)
	//linkedList.Print()
	//fmt.Println("add at tail 3")
	//linkedList.AddAtTail(3)
	//linkedList.Print()
	//fmt.Println("add at index 1 val 2")
	//linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	//linkedList.Print()
	//fmt.Println("get at index 1")
	//fmt.Println("res", linkedList.Get(1)) //返回2
	//linkedList.Print()
	//fmt.Println("delete at index 1")
	//linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	//linkedList.Print()
	//fmt.Println("get at index 1")
	//fmt.Println("res", linkedList.Get(1)) //返回3
	//linkedList.Print()

	//fmt.Println("init")
	//linkedList := ConstructorMyLinkedList()
	//linkedList.Print()
	//fmt.Println("add at head 1")
	//linkedList.AddAtHead(1)
	//linkedList.Print()
	//fmt.Println("add at index 1 val 2")
	//linkedList.AddAtIndex(1, 2)
	//linkedList.Print()
	//fmt.Println("get at index 0")
	//fmt.Println("res", linkedList.Get(1))
	//linkedList.Print()
	//fmt.Println("get at index 1")
	//fmt.Println("res", linkedList.Get(0))
	//linkedList.Print()
	//fmt.Println("get at index 2")
	//fmt.Println("res", linkedList.Get(2))
	//linkedList.Print()

	//fmt.Println("init")
	//linkedList := ConstructorMyLinkedList()
	//linkedList.Print()
	//fmt.Println("add at head 1")
	//linkedList.AddAtHead(1)
	//linkedList.Print()
	//fmt.Println("delete at index 0")
	//linkedList.DeleteAtIndex(0)
	//linkedList.Print()

	//fmt.Println("init")
	//linkedList := ConstructorMyLinkedList()
	//linkedList.Print()
	//fmt.Println("add at index 0 10")
	//linkedList.AddAtIndex(0, 10)
	//linkedList.Print()
	//fmt.Println("add at index 0 20")
	//linkedList.AddAtIndex(0, 20)
	//linkedList.Print()
	//fmt.Println("add at index 1 30")
	//linkedList.AddAtIndex(0, 30)
	//linkedList.Print()
	//fmt.Println("get at index 0")
	//fmt.Println("res", linkedList.Get(0))
	//linkedList.Print()

	fmt.Println("init")
	linkedList := ConstructorMyLinkedList()
	linkedList.Print()
	fmt.Println("add at head 1")
	linkedList.AddAtHead(1)
	linkedList.Print()
	linkedList.AddAtTail(3)
	linkedList.Print()
	fmt.Println("get at index 0")
	fmt.Println("res", linkedList.Get(1))
	linkedList.Print()
	fmt.Println("get at index 1")
	fmt.Println("res", linkedList.Get(0))
	linkedList.Print()
	fmt.Println("get at index 2")
	fmt.Println("res", linkedList.Get(2))
	linkedList.Print()
}

func TestCheckSystemTimer(t *testing.T) {
	CheckSystemTimer()
}

func TestCheckTimerTicker(t *testing.T) {
	CheckTimerTicker()
}

func TestSetTimeoutFunc(t *testing.T) {
	f := func(a int) int {
		fmt.Println(a, "carry func")
		return 12345
	}

	SetTimeoutFunc(TimeTriggerFunc(f), 1*time.Second, 1111111)

	time.Sleep(3 * time.Second)
}

func TestCheckDefer(t *testing.T) {
	CheckDefer()
	assert.Nil(t, nil)
}

func TestShipWithinDays(t *testing.T) {
	assert.Equal(t, ShipWithinDaysV2([]int{1}, 1), 1)
	assert.Equal(t, ShipWithinDaysV2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5), 15)
	assert.Equal(t, ShipWithinDaysV2([]int{1, 2, 3, 1, 1}, 4), 3)
	assert.Equal(t, ShipWithinDaysV2([]int{3, 2, 2, 4, 1, 4}, 3), 6)
	assert.Equal(t, ShipWithinDaysV2([]int{180, 373, 75, 82, 497, 23, 303, 299, 53, 426, 152, 314, 206, 433, 283, 370, 179, 254, 265, 431, 453, 17, 189, 224}, 12), 631)
}

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, IntToRoman(3), "III")
	assert.Equal(t, IntToRoman(4), "IV")
	assert.Equal(t, IntToRoman(9), "IX")
	assert.Equal(t, IntToRoman(58), "LVIII")
	assert.Equal(t, IntToRoman(1994), "MCMXCIV")
}
