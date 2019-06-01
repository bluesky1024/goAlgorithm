package design

import (
	"container/list"
)

/*问题*/
/*
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。
示例:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false
说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/
/*思路*/
/*
栈  ：push to top; peek from top; pop from top;
队列：push to tail；peek from top； pop from top；

两个栈a、b实现：
初始数据固定push进栈a中
要pop或者peek时，先检查b中是否为空
	若为空，则将栈a中所有数据pop出来，push进栈b中
	若不为空，直接将栈b中的数据pop或者peek

*/
type MyQueue struct {
	listA *list.List
	listB *list.List
}

/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{
		listA: list.New(),
		listB: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.listA.PushFront(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.listB.Len() != 0 {
		res := this.listB.Front().Value.(int)
		this.listB.Remove(this.listB.Front())
		return res
	}

	for this.listA.Len() > 0 {
		temp := this.listA.Front().Value.(int)
		this.listA.Remove(this.listA.Front())
		this.listB.PushFront(temp)
	}

	res := this.listB.Front().Value.(int)
	this.listB.Remove(this.listB.Front())
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.listB.Len() != 0 {
		res := this.listB.Front().Value.(int)
		return res
	}

	for this.listA.Len() > 0 {
		temp := this.listA.Front().Value.(int)
		this.listA.Remove(this.listA.Front())
		this.listB.PushFront(temp)
	}

	res := this.listB.Front().Value.(int)
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.listA.Len() == 0 && this.listB.Len() == 0
}

/*问题*/
/*
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
*/
/*思路*/
/*
栈  ：push to top; peek from top; pop from top;
队列：push to tail；peek from top； pop from top；

两个队列a、b实现：
初始数据固定push进有数据的那个队列，初始没数据随意
要pop或者peek时，选择有数据那个队列，每次都把数据从top开始pop到没数据的队列
	若为空，则将栈a中所有数据pop出来，push进栈b中
	若不为空，直接将栈b中的数据pop或者peek
*/
type MyStack struct {
	listA *list.List
	listB *list.List
}

/** Initialize your data structure here. */
func MyStackConstructor() MyStack {
	return MyStack{
		listA: list.New(),
		listB: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.listA.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	NotEmptyList := this.listA
	EmptyList := this.listB
	if this.listA.Len() == 0 {
		NotEmptyList = this.listB
		EmptyList = this.listA
	}

	for NotEmptyList.Len() > 0 {
		temp := NotEmptyList.Front().Value.(int)
		if NotEmptyList.Len() == 1 {
			NotEmptyList.Remove(NotEmptyList.Front())
			return temp
		}
		NotEmptyList.Remove(NotEmptyList.Front())
		EmptyList.PushBack(temp)
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	NotEmptyList := this.listA
	EmptyList := this.listB
	if this.listA.Len() == 0 {
		NotEmptyList = this.listB
		EmptyList = this.listA
	}

	for NotEmptyList.Len() > 0 {
		temp := NotEmptyList.Front().Value.(int)
		if NotEmptyList.Len() == 1 {
			NotEmptyList.Remove(NotEmptyList.Front())
			EmptyList.PushBack(temp)
			return temp
		}
		NotEmptyList.Remove(NotEmptyList.Front())
		EmptyList.PushBack(temp)
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.listA.Len() == 0 && this.listB.Len() == 0
}
