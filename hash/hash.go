package hash

/*问题*/
/*
在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。

返回重复了 N 次的那个元素。



示例 1：

输入：[1,2,3,3]
输出：3
示例 2：

输入：[2,1,2,5,3,2]
输出：2
示例 3：

输入：[5,1,5,2,5,3,5,4]
输出：5


提示：

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length 为偶数
*/
/*思路*/
/*
N+1个元素，有一个元素重复N次
说明其他元素只有一次
说明如果有个元素重复了，那一定就是目标元素
*/
func RepeatedNTimes(A []int) int {
	tempMap := make(map[int]bool, len(A)/2)
	for _, v := range A {
		_, ok := tempMap[v]
		if ok {
			return v
		} else {
			tempMap[v] = true
		}
	}
	return 0
}

/*问题*/
/*
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。

h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”



示例:

输入: citations = [3,0,6,1,5]
输出: 3
解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。


说明: 如果 h 有多种可能的值，h 指数是其中最大的那个。
*/
/*思路*/
/*
至多有h篇文章分别被至少引用了h次（被引用次数为100的文章，也可以看作至少被引用了一次）
*/
func HIndex(citations []int) int {
	tempRes := make(map[int]int, len(citations))
	for _, v1 := range citations {
		if v1 == 0 {
			continue
		}
		temp := 0
		for _, v2 := range citations {
			if v1 <= v2 {
				temp++
			}
		}
		tempRes[v1] = temp
	}

	maxH := 0
	for k, v := range tempRes {
		min := k
		if k > v {
			min = v
		}
		if maxH < min {
			maxH = min
		}
	}
	return maxH
}
