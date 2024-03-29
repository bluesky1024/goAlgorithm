package design

import (
	"container/list"
	"context"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
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
func (q *MyQueue) Push(x int) {
	q.listA.PushFront(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.listB.Len() != 0 {
		res := q.listB.Front().Value.(int)
		q.listB.Remove(q.listB.Front())
		return res
	}

	for q.listA.Len() > 0 {
		temp := q.listA.Front().Value.(int)
		q.listA.Remove(q.listA.Front())
		q.listB.PushFront(temp)
	}

	res := q.listB.Front().Value.(int)
	q.listB.Remove(q.listB.Front())
	return res
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.listB.Len() != 0 {
		res := q.listB.Front().Value.(int)
		return res
	}

	for q.listA.Len() > 0 {
		temp := q.listA.Front().Value.(int)
		q.listA.Remove(q.listA.Front())
		q.listB.PushFront(temp)
	}

	res := q.listB.Front().Value.(int)
	return res
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.listA.Len() == 0 && q.listB.Len() == 0
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
func (s *MyStack) Push(x int) {
	s.listA.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	NotEmptyList := s.listA
	EmptyList := s.listB
	if s.listA.Len() == 0 {
		NotEmptyList = s.listB
		EmptyList = s.listA
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
func (s *MyStack) Top() int {
	NotEmptyList := s.listA
	EmptyList := s.listB
	if s.listA.Len() == 0 {
		NotEmptyList = s.listB
		EmptyList = s.listA
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
func (s *MyStack) Empty() bool {
	return s.listA.Len() == 0 && s.listB.Len() == 0
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
func (h *MyHashMap) Put(key int, value int) {
	ind := key % 10000
	if h.hashMap[ind] == nil {
		h.hashMap[ind] = list.New()
	}

	//遍历整个list，如果存在该key，就更新，不存在就pushback
	for e := h.hashMap[ind].Front(); e != nil; e = e.Next() {
		if e.Value.(SingleHashKV).key == key {
			e.Value = SingleHashKV{
				key: key,
				val: value,
			}
			return
		}
	}
	h.hashMap[ind].PushBack(SingleHashKV{
		key: key,
		val: value,
	})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (h *MyHashMap) Get(key int) int {
	ind := key % 10000
	if h.hashMap[ind] == nil {
		return -1
	}
	for e := h.hashMap[ind].Front(); e != nil; e = e.Next() {
		if e.Value.(SingleHashKV).key == key {
			return e.Value.(SingleHashKV).val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (h *MyHashMap) Remove(key int) {
	ind := key % 10000
	if h.hashMap[ind] == nil {
		return
	}
	for e := h.hashMap[ind].Front(); e != nil; e = e.Next() {
		if e.Value.(SingleHashKV).key == key {
			h.hashMap[ind].Remove(e)
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

func (l *MyLinkedList) Print() {
	if l.nodes == nil {
		fmt.Println("empty list")
	}
	cur := l.nodes
	for {
		if cur == nil {
			break
		}
		fmt.Println(cur.val)
		cur = cur.next
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index+1 > l.len {
		return -1
	}
	cur := l.nodes
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
func (l *MyLinkedList) AddAtHead(val int) {
	l.len++
	if l.nodes == nil {
		l.nodes = &MyLinkedNode{
			val:  val,
			next: nil,
		}
		return
	}
	newNode := &MyLinkedNode{
		val:  val,
		next: l.nodes,
	}
	l.nodes = newNode
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	l.len++
	if l.nodes == nil {
		l.nodes = &MyLinkedNode{
			val:  val,
			next: nil,
		}
		return
	}
	cur := l.nodes
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
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	l.len++
	if l.nodes == nil {
		if index == 0 {
			l.nodes = &MyLinkedNode{
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
	cur := l.nodes
	for {
		if index == 0 {
			if last == nil {
				newNode.next = cur
				l.nodes = newNode
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
func (l *MyLinkedList) DeleteAtIndex(index int) {
	l.len--
	if l.nodes == nil {
		return
	}
	var last *MyLinkedNode = nil
	cur := l.nodes
	for {
		if index == 0 {
			if last == nil {
				l.nodes = nil
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

/*问题*/
/*
设计定时器
*/
/*思路*/
/*
 */

//系统自带定时器了解，便于查看源码
func CheckSystemTimer() {
	input := make(chan interface{})
	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
		input <- "hello, world"
	}()

	t1 := time.NewTimer(5 * time.Second)
	t2 := time.NewTimer(10 * time.Second)

	for {
		select {
		case msg := <-input:
			fmt.Println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(5 * time.Second)

		case <-t2.C:
			println("10s timer")
			t2.Reset(10 * time.Second)
		}
	}
}

func CheckTimerTicker() {
	time1 := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-time1.C:
			fmt.Println("carry finish")
			return
		case a := <-time.After(1 * time.Second):
			fmt.Println("time after", a)
		}
	}
}

//自行实现定时器

type TimeTriggerFunc func(a int) int

func SetTimeoutFunc(triggerFunc TimeTriggerFunc, t time.Duration, a int) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	go func() {
		<-ctx.Done()
		b := triggerFunc(a)
		fmt.Println(b)
	}()
}

func catchPanic() {
	defer func() {
		panic("panic inner")
	}()

	//if err := recover(); err != nil {
	//	fmt.Println("catch panic", err)
	//}
}

func CheckDefer() {
	defer func() {
		fmt.Println("enter 1")
		//if err := recover(); err != nil {
		//	fmt.Println(123, err)
		//}
	}()

	defer func() {
		fmt.Println("enter 2")
		if err := recover(); err != nil {
			fmt.Println(234, err)
		}
	}()

	defer func() {
		fmt.Println("enter 3")
		catchPanic()
	}()

	fmt.Println("aaa")
	panic("panic trigger")
}

/*问题*/
/*
传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。



示例 1：

输入：weights = [1,2,3,4,5,6,7,8,9,10], D = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
示例 2：

输入：weights = [3,2,2,4,1,4], D = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4
示例 3：

输入：weights = [1,2,3,1,1], D = 4
输出：3
解释：
第 1 天：1
第 2 天：2
第 3 天：3
第 4 天：1, 1


提示：

1 <= D <= weights.length <= 50000
1 <= weights[i] <= 500

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
思路一
过于愚蠢
递归+剪枝
按weights的内容今天搬运还是之后搬运分为两波
剪枝方式
如果天数<weight数组，可以舍弃
如果weight数组已经大于其他组最大值，可以舍弃
思路二
结果可选范围为从1到sum(weights)
直接遍历都能完成，但是可以使用二分查找进行优化
*/
func ShipWithinDays(weights []int, D int) int {
	return shipWithinDays(weights, D, 0)
}

func shipWithinDays(weights []int, D int, thisDayOriWeight int) int {
	// weights 空了，直接返回0
	if len(weights) == 0 {
		return thisDayOriWeight
	}

	// 最后一天，只能全部搬运了
	if D == 1 {
		sum := thisDayOriWeight
		for _, weight := range weights {
			sum += weight
		}
		return sum
	}

	// weights[0] 算在今天搬运
	maxWeight1 := shipWithinDays(weights[1:], D, thisDayOriWeight+weights[0])
	if thisDayOriWeight+weights[0] > maxWeight1 {
		maxWeight1 = thisDayOriWeight + weights[0]
	}

	// weights[0] 不算在今天搬运
	maxWeight2 := shipWithinDays(weights, D-1, 0)
	if thisDayOriWeight > maxWeight2 {
		maxWeight2 = thisDayOriWeight
	}

	if maxWeight1 > maxWeight2 {
		return maxWeight2
	}
	return maxWeight1
}

func ShipWithinDaysV2(weights []int, D int) int {
	isMatch := func(weights []int, D int, curData int) bool {
		dayCnt := 1
		curDayWeight := 0
		for _, weight := range weights {
			if curDayWeight+weight > curData {
				dayCnt++
				curDayWeight = weight
				continue
			}
			curDayWeight += weight
		}
		return dayCnt <= D
	}

	min := 0
	max := 0
	for _, weight := range weights {
		if weight > min {
			min = weight
		}
		max += weight
	}
	cur := (min + max) / 2
	for {
		if !isMatch(weights, D, cur) {
			min = cur + 1
		} else {
			max = cur
		}

		if (min+max)/2 == cur {
			break
		}
		cur = (min + max) / 2
	}
	return cur
}

/*问题*/
/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。



示例 1:

输入: 3
输出: "III"
示例 2:

输入: 4
输出: "IV"
示例 3:

输入: 9
输出: "IX"
示例 4:

输入: 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:

输入: 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.


提示：

1 <= num <= 3999

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
首先肯定需要将数据从高进制到低进制取整取模
最大输入是3999，所以顶多是4个X

I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func IntToRoman(num int) string {
	curData := num
	res := ""

	// M 渲染
	numM := curData / 1000
	if numM > 0 {
		res += strings.Repeat("M", numM)
	}
	curData = curData % 1000

	// 100 - 1000 渲染
	num100 := curData / 100
	if num100 == 9 {
		res += "CM"
	} else if num100 == 4 {
		res += "CD"
	} else if num100 >= 5 {
		res += "D"
		res += strings.Repeat("C", (curData%500)/100)
	} else {
		res += strings.Repeat("C", num100)
	}
	curData = curData % 100

	// 10 - 100 渲染
	num10 := curData / 10
	if num10 == 9 {
		res += "XC"
	} else if num10 == 4 {
		res += "XL"
	} else if num10 >= 5 {
		res += "L"
		res += strings.Repeat("X", (curData%50)/10)
	} else {
		res += strings.Repeat("X", num10)
	}
	curData = curData % 10

	// 1 - 10 渲染
	if curData == 9 {
		res += "IX"
	} else if curData == 4 {
		res += "IV"
	} else if curData >= 5 {
		res += "V"
		res += strings.Repeat("I", curData%5)
	} else {
		res += strings.Repeat("I", curData)
	}

	return res
}

/*问题*/
/*
假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N) 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：

第 i 位的数字能被 i 整除
i 能被第 i 位上的数字整除
现在给定一个整数 N，请问可以构造多少个优美的排列？

示例1:

输入: 2
输出: 2
解释:

第 1 个优美的排列是 [1, 2]:
  第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
  第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除

第 2 个优美的排列是 [2, 1]:
  第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
  第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
说明:

N 是一个正整数，并且不会超过15。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/beautiful-arrangement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
/*思路*/
/*
n[i] % i == 0 || i % n[i] == 0
1
1
2
1 2；2 1
3
1 2 3；3 2 1

写了几个可以看出，最好的方法，采用动态规划
求出s[n-1] 再求 s[n] 能在s[n-1]基础上怎么安排num[n]

直接计算cnt比较难递推，还是需要把可能的组合全部列出来

上述想的太简单了，以为s[n]相较s[n-1]只是调整了两个数的位置，其实可能有3个或以上的数据调整位置

最后用了暴力方法，从1到n遍历，逐个推入备选答案中，使用预设条件进行剪枝，错误答案也保留吧，TODO 想想怎么把错误方法改造正确
*/

func CountArrangementWrong(n int) int {
	// 在n-1基础上，n这个数可以存放的位置来自于它的约数
	// n这个数的约数所在位置的数与n互换，首先n到约数位置肯定是没问题的，但约数到n的位置不一定可行，怎么过滤？？？
	// 另外因为在1-n中n肯定是最大值，所以也不需要考虑它的倍数的情况）
	res := queryArrangement(n)
	return len(res)
}

func queryArrangement(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	res := [][]int{{1}}
	for i := 2; i <= n; i++ {
		divisors := getDivisors(i)
		tempRes := make([][]int, 0)
		for _, divisor := range divisors {
			if divisor == i {
				for _, l := range res {
					temp := make([]int, i)
					copy(temp, l)
					temp[i-1] = i
					tempRes = append(tempRes, temp)
				}
				continue
			}
			for _, l := range res {
				if i%l[divisor-1] == 0 {
					temp := make([]int, i)
					copy(temp, l)
					temp[i-1] = temp[divisor-1]
					temp[divisor-1] = i
					tempRes = append(tempRes, temp)
				}
			}
		}

		res = tempRes
	}
	return res
}

func getDivisors(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
		}
	}
	return res
}

func CountArrangement(N int) int {
	res := 0

	var checkArrangeMent func(cur []int, usedMap map[int]struct{})
	checkArrangeMent = func(cur []int, usedMap map[int]struct{}) {
		if len(cur) == N {
			res++
			return
		}
		pos := len(cur) + 1
		for i := 1; i <= N; i++ {
			if _, ok := usedMap[i]; ok {
				continue
			}
			if i%pos == 0 || pos%i == 0 {
				cur = append(cur, i)
				usedMap[i] = struct{}{}
				checkArrangeMent(cur, usedMap)
				delete(usedMap, i)
				cur = cur[:len(cur)-1]
			}
		}
	}

	checkArrangeMent([]int{}, make(map[int]struct{}))

	return res
}

/*问题*/
/*

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/roman-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。



示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.


提示：

1 <= s.length <= 15
s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
IL 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
关于罗马数字的详尽书写规则
*/
/*思路*/
/*
每个数都有自己的数值，当前一个数比后一个数小的时候执行取反操作，逐个相加
*/
func RomanToInt(s string) int {
	dataMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	cur := 0
	last := rune(0)
	for _, r := range s {
		if last == 0 {
			cur += dataMap[r]
		} else if dataMap[last] >= dataMap[r] {
			res += cur
			cur = dataMap[r]
		} else {
			cur = -1*cur + dataMap[r]
		}
		last = r
	}
	res += cur

	return res
}

/*问题*/
/*
给定一个正整数数组 w ，其中 w[i] 代表下标 i 的权重（下标从 0 开始），请写一个函数 pickIndex ，它可以随机地获取下标 i，选取下标 i 的概率与 w[i] 成正比。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。

也就是说，选取下标 i 的概率为 w[i] / sum(w) 。



示例 1：

输入：
["Solution","pickIndex"]
[[[1]],[]]
输出：
[null,0]
解释：
Solution solution = new Solution([1]);
solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
示例 2：

输入：
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出：
[null,1,1,1,1,0]
解释：
Solution solution = new Solution([1, 3]);
solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。

由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
[null,1,1,1,1,0]
[null,1,1,1,1,1]
[null,1,1,1,0,0]
[null,1,1,1,0,1]
[null,1,0,1,0,0]
......
诸若此类。


提示：

1 <= w.length <= 10000
1 <= w[i] <= 10^5
pickIndex 将被调用不超过 10000 次

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/random-pick-with-weight
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
思路相对比较简单，通过random包来模拟随机
权重用数组分段来替代
[2,3,5] 对应的分段是
{0,1} {2,3,4} {5,6,7,8,9}
设置随机数从0到9出现
判断该数据在哪个分段即可pickIdx
判断分段可以用二分法优化一下查询
*/

type Solution struct {
	mapWeight []int64
	sum       int64
}

func Constructor(w []int) Solution {
	mapWeight := make([]int64, 0, len(w))
	sum := int64(0)
	for _, d := range w {
		mapWeight = append(mapWeight, sum)
		sum += int64(d)
	}

	return Solution{
		sum:       sum,
		mapWeight: mapWeight,
	}
}

func (s *Solution) PickIndex() int {
	idx := rand.Int63n(s.sum)
	//idx := int64(5)
	// 二分法查找
	head := 0
	tail := len(s.mapWeight) - 1
	mid := 0
	for head <= tail {
		mid = (head + tail) / 2
		if s.mapWeight[mid] > idx {
			tail = mid - 1
		} else if s.mapWeight[mid] < idx {
			if (mid+1 <= len(s.mapWeight)-1) && (s.mapWeight[mid+1] > idx) {
				break
			}
			if mid+1 > len(s.mapWeight)-1 {
				break
			}
			head = mid + 1
		} else {
			break
		}
	}
	return mid
}

/*问题*/
/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。


进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？



示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4


提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
纯粹参考 group cache
LRU 即最近最少使用
维护一个有序环形链表，每次get或者put操作，都将这个key对应的元素放到列表的头部；
然后当put的时候超过了cap时，就从环形列表的尾部删除最后一个元素
可能需要注意的点，删除元素的时候需要加锁保护
cap 最大才3000，直接初始化map的时候就提供这么大也无所谓
*/
type lruNode struct {
	key int
	val int
	ele *list.Element
}

type LRUCache struct {
	cap    int
	data   map[int]*lruNode
	orderL *list.List
	mutex  sync.RWMutex
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{
		data:   make(map[int]*lruNode, capacity),
		orderL: list.New(),
		cap:    capacity,
		mutex:  sync.RWMutex{},
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.data[key]; ok {
		l.orderL.MoveToFront(node.ele)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	// 查询key在存储中是否存在
	if node, ok := l.data[key]; ok {
		l.orderL.MoveToFront(node.ele)
		node.val = value
		return
	}

	// 不存在
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// cap 满了，先删除再插入
	if len(l.data) == l.cap {
		l.delBackEle()
	}

	// 插入新的key
	node := &lruNode{
		key: key,
		val: value,
	}
	node.ele = l.orderL.PushFront(node)
	l.data[key] = node
}

func (l *LRUCache) delBackEle() {
	back := l.orderL.Back()
	if back == nil {
		return
	}
	node := back.Value.(*lruNode)
	delete(l.data, node.key)
	l.orderL.Remove(back)
}

/*

题目描述
假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。
注意：
总人数少于1100人。

示例
输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

先从小到大排序，从4开始找位置
44 50 52 61 70 71

_ _ _ _ _ _
5 7 5 6 4 7

*/

type mPair struct {
	h int
	k int
}

func rebuild(pairs []mPair) []mPair {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].h < pairs[j].h {
			return true
		}
		if pairs[i].h > pairs[j].h {
			return false
		}
		return pairs[i].k < pairs[j].k
	})
	res := make([]mPair, len(pairs))
	for _, p := range pairs {
		cur := 0
		// 找到p的k对应的位置
		for i, pp := range res {
			if pp.h == 0 {
				if cur == p.k {
					res[i] = p
					break
				}
				cur++
				continue
			}

			if pp.h >= p.h {
				cur++
			}
		}
	}
	return res
}
