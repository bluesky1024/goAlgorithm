package design

import (
	"container/list"
	"fmt"
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

/*问题*/
/*
不使用任何内建的哈希表库设计一个哈希映射

具体地说，你的设计应该包含以下的功能

put(key, value)：向哈希映射中插入(键,值)的数值对。如果键对应的值已经存在，更新这个值。
get(key)：返回给定的键所对应的值，如果映射中不包含这个键，返回-1。
remove(key)：如果映射中存在这个键，删除这个数值对。

示例：

MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // 返回 1
hashMap.get(3);            // 返回 -1 (未找到)
hashMap.put(2, 1);         // 更新已有的值
hashMap.get(2);            // 返回 1
hashMap.remove(2);         // 删除键为2的数据
hashMap.get(2);            // 返回 -1 (未找到)

注意：

所有的值都在 [1, 1000000]的范围内。
操作的总数目在[1, 10000]范围内。
不要使用内建的哈希库。
*/
/*思路*/
/*
1
直观来看，首先想到的是取模，数据存在则往后顺延
因为操作的总数<=1w，直接建立1w的数组，初始化为0，如果出现碰撞则往后顺延，找到第一个不为0的值
这种方案第一次提交出现问题，在于remove的for循环退出的条件不对，如果去除那个==0的条件，又会导致删除操作每次都要遍历整个map，效率简直低

2
建立二维数组吧
第一层根据key进行hash
第二层建立一个list
*/
type SingleHashKV struct {
	key int
	val int
}
type MyHashMap struct {
	hashMap []*list.List
}

/** Initialize your data structure here. */
func MyHashMapConstructor() MyHashMap {
	return MyHashMap{
		hashMap: make([]*list.List, 10000),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	ind := key % 10000
	if this.hashMap[ind] == nil {
		this.hashMap[ind] = list.New()
	}

	//遍历整个list，如果存在该key，就更新，不存在就pushback
	for e := this.hashMap[ind].Front(); e != nil; e = e.Next() {
		if e.Value.(SingleHashKV).key == key {
			e.Value = SingleHashKV{
				key: key,
				val: value,
			}
			return
		}
	}
	this.hashMap[ind].PushBack(SingleHashKV{
		key: key,
		val: value,
	})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	ind := key % 10000
	if this.hashMap[ind] == nil {
		return -1
	}
	for e := this.hashMap[ind].Front(); e != nil; e = e.Next() {
		if e.Value.(SingleHashKV).key == key {
			return e.Value.(SingleHashKV).val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	ind := key % 10000
	if this.hashMap[ind] == nil {
		return
	}
	for e := this.hashMap[ind].Front(); e != nil; e = e.Next() {
		if e.Value.(SingleHashKV).key == key {
			this.hashMap[ind].Remove(e)
		}
	}
}

/*问题*/
/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
初始化的时候有一个节点，此时val==-1表示这个节点是无效的，即list是空list
*/
type MyLinkedList struct {
	len   int
	nodes *MyLinkedNode
}

type MyLinkedNode struct {
	val  int
	next *MyLinkedNode
}

/** Initialize your data structure here. */
func ConstructorMyLinkedList() MyLinkedList {
	return MyLinkedList{
		nodes: nil,
		len:   0,
	}
}

func (this *MyLinkedList) Print() {
	if this.nodes == nil {
		fmt.Println("empty list")
	}
	cur := this.nodes
	for {
		if cur == nil {
			break
		}
		fmt.Println(cur.val)
		cur = cur.next
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index+1 > this.len {
		return -1
	}
	cur := this.nodes
	for {
		if cur == nil {
			break
		}
		if index == 0 {
			return cur.val
		}
		index--
		cur = cur.next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.len++
	if this.nodes == nil {
		this.nodes = &MyLinkedNode{
			val:  val,
			next: nil,
		}
		return
	}
	newNode := &MyLinkedNode{
		val:  val,
		next: this.nodes,
	}
	this.nodes = newNode
	return
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.len++
	if this.nodes == nil {
		this.nodes = &MyLinkedNode{
			val:  val,
			next: nil,
		}
		return
	}
	cur := this.nodes
	for {
		if cur.next == nil {
			cur.next = &MyLinkedNode{
				val:  val,
				next: nil,
			}
			return
		}
		cur = cur.next
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	this.len++
	if this.nodes == nil {
		if index == 0 {
			this.nodes = &MyLinkedNode{
				val:  val,
				next: nil,
			}
		}
		return
	}
	newNode := &MyLinkedNode{
		val:  val,
		next: nil,
	}
	var last *MyLinkedNode = nil
	cur := this.nodes
	for {
		if index == 0 {
			if last == nil {
				newNode.next = cur
				this.nodes = newNode
			} else {
				newNode.next = cur
				last.next = newNode
			}
			return
		}
		index--
		last = cur
		cur = cur.next
		if cur == nil {
			if index == 0 {
				last.next = newNode
			}
			return
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	this.len--
	if this.nodes == nil {
		return
	}
	var last *MyLinkedNode = nil
	cur := this.nodes
	for {
		if index == 0 {
			if last == nil {
				this.nodes = nil
			} else {
				last.next = cur.next
			}
			return
		}
		index--
		last = cur
		cur = cur.next
		if cur == nil {
			return
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
