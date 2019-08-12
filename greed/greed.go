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

/*问题*/
/*
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

注意：

你可以假设胃口值为正。
一个小朋友最多只能拥有一块饼干。

示例 1:

输入: [1,2,3], [1,1]

输出: 1

解释:
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1。
示例 2:

输入: [1,2], [1,2,3]

输出: 2

解释:
你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/assign-cookies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
饼干尺寸和小孩胃口从小到大排列
从胃口最小的孩子开始满足，满足的时候尽量用小点的饼干
*/
func FindContentChildren(g []int, s []int) int {
	lengthG := len(g)
	if lengthG == 0 {
		return 0
	}

	oriSort.Ints(g)
	oriSort.Ints(s)

	res := 0
	gInd := 0
	for _, sI := range s {
		for {
			if g[gInd] >= sI {
				gInd++
				res++
				break
			}
			gInd++
			if gInd >= lengthG {
				break
			}
		}

		if gInd >= lengthG {
			break
		}
	}

	return res
}

/*问题*/
/*
假设你是一位顺风车司机，车上最初有 capacity 个空座位可以用来载客。由于道路的限制，车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向，你可以将其想象为一个向量）。

这儿有一份行程计划表 trips[][]，其中 trips[i] = [num_passengers, start_location, end_location] 包含了你的第 i 次行程信息：

必须接送的乘客数量；
乘客的上车地点；
以及乘客的下车地点。
这些给出的地点位置是从你的 初始 出发位置向前行驶到这些地点所需的距离（它们一定在你的行驶方向上）。

请你根据给出的行程计划表和车子的座位数，来判断你的车是否可以顺利完成接送所用乘客的任务（当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false）。



示例 1：

输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
示例 2：

输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true
示例 3：

输入：trips = [[2,1,5],[3,5,7]], capacity = 3
输出：true
示例 4：

输入：trips = [[3,2,7],[3,7,9],[8,3,9]], capacity = 11
输出：true


提示：

你可以假设乘客会自觉遵守 “先下后上” 的良好素质
trips.length <= 1000
trips[i].length == 3
1 <= trips[i][0] <= 100
0 <= trips[i][1] < trips[i][2] <= 1000
1 <= capacity <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/car-pooling
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1.按上车位置进行排序
2.逐点进行位置更新
或者
1000个点，每个点处记录人数变化
从0开始遍历到1000，中途哪里超员即可返回false，否则返回true
*/
type CarPoolPara struct {
	para [][]int
}

// Len is the number of elements in the collection.
func (c CarPoolPara) Len() int {
	return len(c.para)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c CarPoolPara) Less(i, j int) bool {
	return c.para[i][1] < c.para[j][1]
}

// Swap swaps the elements with indexes i and j.
func (c CarPoolPara) Swap(i, j int) {
	c.para[i], c.para[j] = c.para[j], c.para[i]
}

func CarPooling(trips [][]int, capacity int) bool {
	pools := CarPoolPara{
		para: trips,
	}
	oriSort.Sort(pools)

	curNum := 0
	curPos := -1
	for _, trip := range pools.para {
		curPos = trip[1]
		for j, tripSub := range pools.para {
			if tripSub[0] > 0 && tripSub[2] <= curPos {
				curNum = curNum - tripSub[0]
				pools.para[j][0] = 0
			}
		}
		curNum = curNum + trip[0]
		if curNum > capacity {
			return false
		}
	}
	return true
}
