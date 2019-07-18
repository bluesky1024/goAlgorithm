package greed

import (
	"container/list"
	"github.com/bluesky1024/goAlgorithm/sort"
	oriSort "sort"
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

/*问题*/
/*
公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。

返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。



示例：

输入：[[10,20],[30,200],[400,50],[30,20]]
输出：110
解释：
第一个人去 A 市，费用为 10。
第二个人去 A 市，费用为 30。
第三个人去 B 市，费用为 50。
第四个人去 B 市，费用为 20。

最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。


提示：

1 <= costs.length <= 100
costs.length 为偶数
1 <= costs[i][0], costs[i][1] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-city-scheduling
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
总共2N人，2座城市每个城市都有N人抵达
也就是每个人只能二选一，不能不选
按差价排序，按差价高的人先选
优先选价格低的城市
如果城市人数够了，则只能选另一座城市
*/
type peopleDiff struct {
	ind  int
	diff int
}
type diffList []*peopleDiff

func (p diffList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p diffList) Len() int           { return len(p) }
func (p diffList) Less(i, j int) bool { return p[i].diff > p[j].diff }
func TwoCitySchedCost(costs [][]int) int {
	peopleCnt := len(costs)
	cityCnt := peopleCnt / 2
	cityACnt := 0
	cityBCnt := 0

	diffArr := diffList{}
	for i, v := range costs {
		tempDiff := v[1] - v[0]
		if tempDiff < 0 {
			tempDiff = tempDiff * -1
		}
		temp := &peopleDiff{
			ind:  i,
			diff: tempDiff,
		}
		diffArr = append(diffArr, temp)
	}
	oriSort.Sort(diffArr)

	res := 0
	for _, tempDiff := range diffArr {
		if costs[tempDiff.ind][0] < costs[tempDiff.ind][1] {
			if cityACnt < cityCnt {
				cityACnt++
				res = res + costs[tempDiff.ind][0]
			} else {
				cityBCnt++
				res = res + costs[tempDiff.ind][1]
			}
		} else {
			if cityBCnt < cityCnt {
				cityBCnt++
				res = res + costs[tempDiff.ind][1]
			} else {
				cityACnt++
				res = res + costs[tempDiff.ind][0]
			}
		}
	}
	return res
}

/*问题*/
/*
给定一个由 '(' 和 ')' 括号组成的字符串 S，我们需要添加最少的括号（ '(' 或是 ')'，可以在任何位置），以使得到的括号字符串有效。

从形式上讲，只有满足下面几点之一，括号字符串才是有效的：

它是一个空字符串，或者
它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
它可以被写作 (A)，其中 A 是有效字符串。
给定一个括号字符串，返回为使结果字符串有效而必须添加的最少括号数。



示例 1：

输入："())"
输出：1
示例 2：

输入："((("
输出：3
示例 3：

输入："()"
输出：0
示例 4：

输入："()))(("
输出：4


提示：

S.length <= 1000
S 只包含 '(' 和 ')' 字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
还是用栈，尽量匹配
左括号入栈，右括号出栈
发现不能出栈的时候，补1个左括号，res++
最后发现栈内还有n个剩余，补n个右括号，res+n

发现内存消耗比别人要多
别人没用栈，只存了当前没出栈的左括号的个数，去除了出栈入栈的动作
所以内存消耗变大很多，差评
*/
func MinAddToMakeValid(S string) int {
	res := 0
	left := list.New()
	for i, ch := range S {
		if ch == '(' {
			left.PushFront(i)
		} else {
			if left.Len() != 0 {
				left.Remove(left.Front())
			} else {
				res++
			}
		}
	}
	res = res + left.Len()
	return res
}
