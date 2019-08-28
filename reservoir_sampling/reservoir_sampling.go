package reservoir_sampling

import (
	"github.com/bluesky1024/goAlgorithm/linked_list"
	"math/rand"
)

/*问题*/
/*
给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

示例:

// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-random-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
典型的蓄水池抽样算法
从一个长度未知的数组中等概率的提取m个子数据
流式数据，长度未知，维护一个数量为m的池子；
后续流式数据进入的时候，第i个数据按照m/i的概率进行选入池中；
池中原有的m个数据中随机挑一个被挤出池中

推导过程：数学归纳法，i时刻，前i个数据中任意一个出现在池中的概率是m/i。
已有m个数据，当时刻i=m+1的时候，第i个数据进入池中的概率为m/m+1.
已有的m个数据，被替换出去的概率是m/m+1 * 1/m = 1/m+1; 所以该m个数据在时刻i还在池中的概率是m/m+1

假设j=i+1时刻前，上述假设成立，即m中原有的所有数据出现在池中的均为m/i；
第j=i+1个数据是否被选入池中的概率为m/i+1,
原有的m个数据，在j=i+1时刻被替换的概率是 (m/i+1) * (1/m) = 1/(i+1),
不被替换的概率是 1-1/(1+i) = i/(i+1)
基于前i条数据出现于m池中的概率都为 m/i,则任意一条数据在第j=i+1时刻还在池中的概率为 m/i * i/(i+1) = m/(i+1)
所以递归条件成立，前i+1条数据出现在池中的概率都为 m/(i+1)
*/

type Solution struct {
	head *linked_list.ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func SolutionConstructor(head *linked_list.ListNode) Solution {
	return Solution{
		head: head,
	}
}

/** Returns a random node's value. */
func (t *Solution) GetRandom() int {
	if t.head == nil {
		return 0
	}

	res := t.head.Val
	cur := t.head
	ind := 1
	for cur.Next != nil {
		cur = cur.Next
		if rand.Intn(ind+1) == ind {
			res = cur.Val
		}
		ind++
	}

	return res
}
