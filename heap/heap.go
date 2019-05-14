package heap

import (
	"container/heap"
	"math"
)

/*
基础：
顺序存储的完全二叉树可以用数组表示
如下所示：
	     1
	    / \
	   2   3
	  / \
	 4   5
对应的数组即为[1,2,3,4,5]
该数组满足一下规则，若根节点处索引为0，则
1.索引为i的left索引是(2*i)
2.索引为i的right索引是(2*i+1)
3.索引为i的父节点索引是floor(i/2)

最大堆：父结点的键值总是大于或等于任何一个子节点的键值
最小堆：父结点的键值总是小于或等于任何一个子节点的键值

最大堆维护方法：
新增：
1.将新增数据置于数组尾部
2.将该数据与父节点比较，比父节点大则两者互换
3.重复执行，直到小于等于父节点
删除：
1.将数组尾部的数字挪至被删除的空白处
2.
*/

/*问题*/
/*
设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。

你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素。

示例:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
说明:
你可以假设 nums 的长度≥ k-1 且k ≥ 1。
*/
/*思路*/
/*
实时求第K大的元素
维护一个最小堆，堆的大小是K
每次新增一个元素，与堆根元素比较：
若比堆根小，则丢弃
若比堆根大，则将堆根pop，并将该元素push进堆，然后调整成正常最小堆
*/
type KthLargest struct {
	temp  heap.Interface
	k     int
	len   int
	kHeap []int
}

func Constructor(k int, nums []int) KthLargest {
	res := KthLargest{
		k:     k,
		kHeap: make([]int, k),
	}
	for _, v := range nums {
		res.push(v)
	}
	return res
}

func (k *KthLargest) pop() {

}

func (k *KthLargest) push(v int) {
	if k.len < k.k {
		ind := k.len
		k.kHeap[k.len] = v
		for {
			indP := int(math.Floor(float64(ind / 2)))
			if k.kHeap[ind] < k.kHeap[indP] {
				k.kHeap[ind], k.kHeap[indP] = k.kHeap[indP], k.kHeap[ind]
				ind = indP
			}
			if indP == 0 {
				break
			}
		}

		return
	}

	//比较第0个和v的大小
	if k.kHeap[0] >= v {
		return
	}

	//将第0个数换成新push的数，从上往下调整
	k.kHeap[0] = v
	ind := 0
	for {
		indL := 2 * ind
		indR := 2*ind + 1
		if indL < k.len {
			if indR < k.len {
				if v > k.kHeap[indL] {
					k.kHeap[indL], v = v, k.kHeap[indL]
					ind = indL
					continue
				}
				if v > k.kHeap[indR] {
					k.kHeap[indR], v = v, k.kHeap[indR]
					ind = indR
					continue
				}
			} else {
				if v > k.kHeap[indL] {
					k.kHeap[indL], v = v, k.kHeap[indL]
					ind = indL
					continue
				}
			}
		}
		break
	}
	return
}

func (this *KthLargest) Add(v int) int {
	this.push(v)
	return this.kHeap[0]
}
