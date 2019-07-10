package greed

import (
	"github.com/bluesky1024/goAlgorithm/sort"
)

/*问题*/
/*
给定由 N 个小写字母字符串组成的数组 A，其中每个字符串长度相等。

删除 操作的定义是：选出一组要删掉的列，删去 A 中对应列中的所有字符，形式上，第 n 列为 [A[0][n], A[1][n], ..., A[A.length-1][n]]）。

比如，有 A = ["abcdef", "uvwxyz"]，



要删掉的列为 {0, 2, 3}，删除后 A 为["bef", "vyz"]， A 的列分别为["b","v"], ["e","y"], ["f","z"]。



你需要选出一组要删掉的列 D，对 A 执行删除操作，使 A 中剩余的每一列都是 非降序 排列的，然后请你返回 D.length 的最小可能值。



示例 1：

输入：["cba", "daf", "ghi"]
输出：1
解释：
当选择 D = {1}，删除后 A 的列为：["c","d","g"] 和 ["a","f","i"]，均为非降序排列。
若选择 D = {}，那么 A 的列 ["b","a","h"] 就不是非降序排列了。
示例 2：

输入：["a", "b"]
输出：0
解释：D = {}
示例 3：

输入：["zyx", "wvu", "tsr"]
输出：3
解释：D = {0, 1, 2}


提示：

1 <= A.length <= 100
1 <= A[i].length <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-columns-to-make-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
看懂题目就行了，A的每个字符串的长度都一样
所以不存在删除某列后，后段字符串列错位的情况
逐位比较就行了
*/
func MinDeletionSize(A []string) int {
	num := len(A)
	if num <= 1 {
		return 0
	}
	res := 0
	length := len(A[0])
	for indCh := 0; indCh < length; indCh++ {
		for indStr := 0; indStr < num-1; indStr++ {
			if A[indStr][indCh] > A[indStr+1][indCh] {
				res++
				break
			}
		}
	}
	return res
}

/*问题*/
/*
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。



提示：

1 <= stones.length <= 30
1 <= stones[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-stone-weight
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
排序取前两个数，然后置0
余数插入排序
*/
func LastStoneWeight(stones []int) int {
	length := len(stones)
	if length == 1 {
		return stones[0]
	}

	//从大到小
	sort.QuickSort(stones)

	for ind := 0; ind < length-1; ind++ {
		//前两块碰撞
		left := stones[ind] - stones[ind+1]
		stones[ind] = 0
		stones[ind+1] = left

		//重新排序（left挪到合适位置）
		for sortInd := ind + 1; sortInd < length-1; sortInd++ {
			if stones[sortInd] < stones[sortInd+1] {
				stones[sortInd], stones[sortInd+1] = stones[sortInd+1], stones[sortInd]
			} else {
				break
			}
		}
	}

	return stones[length-1]
}
