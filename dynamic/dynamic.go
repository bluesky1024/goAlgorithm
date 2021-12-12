package dynamic

import (
	"fmt"
	"math"
	"sort"
)

/*题目*/
/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
说明:

你可以假设数组不可变。
会多次调用 sumRange 方法。
*/

/*思路*/
/*
以下Node其实可以不要，这个结构体只需要提供SumRange功能，思维僵化了...
*/
type NumArray struct {
	SumInd []int
	Node   []int
}

func ConstructorNumArray(nums []int) NumArray {
	tempSum := 0
	res := NumArray{}
	res.SumInd = make([]int, len(nums))
	res.Node = make([]int, len(nums))
	for ind, v := range nums {
		res.Node[ind] = v
		tempSum = tempSum + v
		res.SumInd[ind] = tempSum
	}
	return res
}

func (a *NumArray) SumRange(i int, j int) int {
	return a.SumInd[j] - a.SumInd[i] + a.Node[i]
}

/*问题*/
/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
[1,3,1],
[1,5,1],
[4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

/*思路*/
/*
根据限制：每次只能向下或者向右移动一步
位置[i,j]的上一步只能是来自于它的上方或者左方（假设上方和左方存在,不存在就忽略）
min[i,j] = min(min[i-1,j],min[i,j-1]) + v[i,j]
*/
func getMinLastStep(minSumArr [][]int, row int, col int) int {
	if row == 0 {
		if col == 0 {
			return 0
		}
		return minSumArr[row][col-1]
	}
	if col == 0 {
		return minSumArr[row-1][col]
	}

	if minSumArr[row-1][col] <= minSumArr[row][col-1] {
		return minSumArr[row-1][col]
	}
	return minSumArr[row][col-1]
}

func MinPathSum(grid [][]int) int {
	minSumArr := make([][]int, len(grid))
	for rowInd, arr := range grid {
		minSumArr[rowInd] = make([]int, len(arr))
		for colInd, v := range arr {
			minSumArr[rowInd][colInd] = getMinLastStep(minSumArr, rowInd, colInd) + v
		}
	}
	return minSumArr[len(minSumArr)-1][len(minSumArr[0])-1]
}

/*问题*/
/*
编写一个程序，找出第 n 个丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
*/
/*思路*/
/*
1.除了1，下一个丑数，一定是之前的丑数*2或者再之前的丑数*3，或者更之前的丑数*5
2.那么问题来了，怎么判断是多少个之前？？？
*/
func NthUglyNumber(n int) int {
	n2 := 0
	n3 := 0
	n5 := 0
	i := 1
	res := make([]int, 1)
	res[0] = 1

	for i = 1; i < n; i++ {
		temp2 := res[n2] * 2
		temp3 := res[n3] * 3
		temp5 := res[n5] * 5

		fmt.Println(res[n2], temp2, res[n3], temp3, res[n5], temp5)

		temp := int(math.Min(float64(temp5), math.Min(float64(temp2), float64((temp3)))))
		res = append(res, temp)
		if temp == temp2 {
			n2++
		}
		if temp == temp3 {
			n3++
		}
		if temp == temp5 {
			n5++
		}
	}
	return res[len(res)-1]
}

/*问题*/
/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
没其他方法，只能遍历，逐层记录到达当前点时的最小路径和
只使用 O(n) 的额外空间，也就是每次循环，要覆盖上一轮的值
覆盖的时候要保证不会再用到，所以应该从后往前更新
*/
func MinimumTotal(triangle [][]int) int {
	levelCnt := len(triangle)
	if levelCnt == 0 {
		return 0
	}
	nodeSum := make([]int, levelCnt)

	for i, nodes := range triangle {
		if i == 0 {
			nodeSum[0] = nodes[0]
			continue
		}

		for ii := i; ii >= 0; ii-- {
			if ii == i {
				nodeSum[ii] = nodeSum[ii-1] + nodes[ii]
			} else if ii == 0 {
				nodeSum[ii] = nodeSum[ii] + nodes[ii]
			} else {
				temp := nodeSum[ii-1]
				if temp > nodeSum[ii] {
					temp = nodeSum[ii]
				}
				nodeSum[ii] = nodes[ii] + temp
			}
		}
	}

	res := 0
	for i, v := range nodeSum {
		if i == 0 {
			res = v
			continue
		}
		if v < res {
			res = v
		}
	}
	return res
}

/*问题*/
/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回符合要求的最少分割次数。

示例:

输入: "aab"
输出: 1
解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
思路a：
1.每个字符之间都能选择是否插入,
2.两种选择都进行判断，并逐步剪支
3.遍历到最后一个字符，然后返回最终的，插入次数
lastCutInd --- 上一个字符串的最后一个字符位置
curInd --- 当前需要判断是否在后方插入，的字符的ind
*/
func checkPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length; i++ {
		if i >= length-i-1 {
			break
		}
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func minCutWithCutTimes(s string, lastCutInd int, curInd int, cutTime int) int {
	if len(s) == curInd+1 {
		if !checkPalindrome(s[lastCutInd+1:]) {
			return 0
		}
		return cutTime
	}
	//若在curInd后方加入，
	s1 := s[lastCutInd+1 : curInd+1]
	if checkPalindrome(s1) {
		cutTime1 := minCutWithCutTimes(s, curInd, curInd+1, cutTime+1)
		cutTime2 := minCutWithCutTimes(s, lastCutInd, curInd+1, cutTime)
		if cutTime1 == 0 && cutTime2 != 0 {
			return cutTime2
		} else if cutTime1 != 0 && cutTime2 == 0 {
			return cutTime1
		} else if cutTime1 == 0 && cutTime2 == 0 {
			return 0
		} else {
			if cutTime1 < cutTime2 {
				return cutTime1
			} else {
				return cutTime2
			}
		}
	} else {
		return minCutWithCutTimes(s, lastCutInd, curInd+1, cutTime)
	}

}

func MinCut(s string) int {
	if checkPalindrome(s) {
		return 0
	}
	curInd := 0
	lastCutInd := -1
	cutTime := 0
	return minCutWithCutTimes(s, lastCutInd, curInd, cutTime)
}

/*问题*/
/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
直观来看，首先肯定从头往后遍历，首先得保证前缀能组成，然后往后递推
为了及时剪枝，需要记录从第n位开始拆分，无法组成
*/
func WordBreak(s string, wordDict []string) bool {
	checkFailInd := make([]bool, len(s))
	res, _ := checkWordBreak(s, wordDict, 0, checkFailInd)
	return res
}

func checkWordBreak(s string, wordDict []string, startInd int, checkFailInd []bool) (bool, int) {
	lengthTarget := len(s)
	if startInd == lengthTarget {
		return true, -1
	}
	if checkFailInd[startInd] {
		return false, startInd
	}
	for _, word := range wordDict {
		lengthWord := len(word)
		if lengthWord+startInd > lengthTarget {
			continue
		}
		if s[startInd:startInd+lengthWord] == word {
			tmpRes, failInd := checkWordBreak(s, wordDict, startInd+lengthWord, checkFailInd)
			if !tmpRes {
				checkFailInd[failInd] = true
			} else {
				return true, -1
			}
		}
	}
	return false, startInd
}

/*问题*/
/*
你正在安装一个广告牌，并希望它高度最大。这块广告牌将有两个钢制支架，两边各一个。每个钢支架的高度必须相等。

你有一堆可以焊接在一起的钢筋 rods。举个例子，如果钢筋的长度为 1、2 和 3，则可以将它们焊接在一起形成长度为 6 的支架。

返回广告牌的最大可能安装高度。如果没法安装广告牌，请返回 0。



示例 1：

输入：[1,2,3,6]
输出：6
解释：我们有两个不相交的子集 {1,2,3} 和 {6}，它们具有相同的和 sum = 6。
示例 2：

输入：[1,2,3,4,5,6]
输出：10
解释：我们有两个不相交的子集 {2,3,5} 和 {4,6}，它们具有相同的和 sum = 10。
示例 3：

输入：[1,2]
输出：0
解释：没法安装广告牌，所以返回 0。


提示：

0 <= rods.length <= 20
1 <= rods[i] <= 1000
钢筋的长度总和最多为 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/tallest-billboard
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
从全集里找出两个子集，使得他们的和相等，并且需要是所有可能中的最大值
钢筋长度总和最多为5000，所以一边的长度至多为2500，超过就不可能两边长度相等
元素长度可以先排序？优先组合最长的结果？
N元素集合，假设已经找到最优结果，引入一个新的元素，所有平衡都被打破？
每个元素有三种选择方式，a.left；b.right;c.抛弃
组合方式 0<=3^n<=3^20=10^10... 有点大,再考虑一些剪支之类的操作，排序是否有必要？

1  2  3  4  5
1  3  6  10 15


5  4  3  2  1
15 10 6  3  1

补充一下剪枝思路：
1.如果left 和 right的差值已经用以下所有的sum都弥补不了，就没必要继续往下递归遍历了
2.如果left + right + 剩下的所有数据的sum都不能比当前最大高度大，也没必要哦继续往下递归遍历了
3.按照从大到小顺序排列，比较容易尽快剪枝
*/

func TallestBillboard(rods []int) int {
	if len(rods) <= 0 {
		return 0
	}

	//先排序
	sort.Ints(rods)
	//sort.Reverse(sort.IntSlice(rods))
	rodSum := make([]int, len(rods))
	for ind := range rods {
		if ind == 0 {
			rodSum[ind] = rods[ind]
			continue
		} else {
			rodSum[ind] = rods[ind] + rodSum[ind-1]
		}
	}
	reverseArr(rodSum)
	reverseArr(rods)

	left := 0
	right := 0
	max := 0
	tallestBillboardWithLeftRightSet(rodSum, rods, left, right, &max)

	return max
}

func reverseArr(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}
}

func tallestBillboardWithLeftRightSet(rodSum []int, rods []int, left int, right int, curMax *int) {
	if len(rods) == 0 {
		return
	}

	//如果所有数据都用上，还是没有当前最大值大，就没必要循环了
	if left+right+rodSum[0] < 2**curMax {
		return
	}

	//放进左边集合
	if left+rods[0] == right && right > *curMax {
		*curMax = right
	}
	if right <= 2500 && left+rods[0] <= 2500 && int(math.Abs(float64(left+rods[0]-right))) <= rodSum[0] {
		tallestBillboardWithLeftRightSet(rodSum[1:], rods[1:], left+rods[0], right, curMax)
	}

	//放进右边集合
	if right+rods[0] == left && left > *curMax {
		*curMax = left
	}
	if left <= 2500 && right+rods[0] <= 2500 && int(math.Abs(float64(right+rods[0]-left))) <= rodSum[0] {
		tallestBillboardWithLeftRightSet(rodSum[1:], rods[1:], left, right+rods[0], curMax)
	}

	//抛弃
	if right == left && left > *curMax {
		*curMax = left
	}
	if left <= 2500 && right <= 2500 && int(math.Abs(float64(right-left))) <= rodSum[0] {
		tallestBillboardWithLeftRightSet(rodSum[1:], rods[1:], left, right, curMax)
	}
}

/*问题*/
/*
给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。

示例:

输入: 2
输出: 91
解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-numbers-with-unique-digits
*/
/*思路*/
/*
y(n) = y(n-1) + x(n)
x(n)表示 10^(n-1)<=x<10^n 中各位数字都不重复的数字个数
x(n)限定了左数第一位不能是0，有1-9 9种可能，剩下n位，因为要求都不能重复，所以从剩下的1-9种剩下的数字和0共9位种挑选n-1个数字随机排序
也就是 9 * A[9 n-1]

另外，如果超过10位，那0-9这10位数字明显不够用了，往上扩展肯定找不到新的符合条件的数字了，直接调整到CountNumbersWithUniqueDigits(10)
*/
func CountNumbersWithUniqueDigits(n int) int {
	if n > 10 {
		return CountNumbersWithUniqueDigits(10)
	}
	if n == 0 {
		return 1
	}

	sum := 1
	for i := 9; i > (10 - n); i-- {
		sum = sum * i
	}
	return CountNumbersWithUniqueDigits(n-1) + sum*9
}

/*问题*/
/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？



示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


限制：

0 <= 数组长度 <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof
*/
/*思路*/
/*
初始化，第一天买入卖出

每到一个点
1.跟之前的卖出值比较,若比之前的卖出值还大，则可以改成今天卖出，更新卖出值和最大收益
2.跟之前的买入值比较，若比之前的买入值还小，则可以更新买入值，以后的交易都可以基于这个买入再进行
3.重新计算最大收益(因为最小买入值可能已经更新了)，若比之前的最大收益还大，则可以更新当前买入值和卖出值和最大收益

直到最后一天
*/

func MaxProfit(prices []int) int {
	length := len(prices)

	if length <= 1 {
		return 0
	}

	curMax := 0
	lastBuyPrice := prices[0]
	lastSellPrice := prices[0]
	for i := 1; i < length; i++ {
		if prices[i] > lastSellPrice {
			lastSellPrice = prices[i]
			curMax = lastSellPrice - lastBuyPrice
			continue
		}

		if prices[i] < lastBuyPrice {
			lastBuyPrice = prices[i]
			continue
		}

		if prices[i]-lastBuyPrice > curMax {
			curMax = prices[i] - lastBuyPrice
			lastSellPrice = prices[i]
		}
	}

	return curMax
}

/*
问题
*/
/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。

图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。



示例:

输入: [2,1,5,6,2,3]
输出: 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
思路
*/
/*
考虑以当前矩形为高，向左向右计算可行的最长宽度
原遍历方式过于简单，导致时间复杂度过高
官方答案中采用了 单调栈 来提高性能，以下直接抄写了官方代码
*/

func LargestRectangleArea(heights []int) int {
	curMax := 0
	for ind, h := range heights {
		left := ind
		for {
			if left == 0 {
				break
			}
			if heights[left-1] >= h {
				left--
				continue
			}
			break
		}
		right := ind
		for {
			if right == len(heights)-1 {
				break
			}
			if heights[right+1] >= h {
				right++
				continue
			}
			break
		}
		length := right - left + 1
		if h*length > curMax {
			curMax = h * length
		}
	}
	return curMax
}

// 官方答案
func LargestRectangleAreaFormal(heights []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	monoStack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

/*问题*/
/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。


示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:

输入: numRows = 1
输出: [[1]]


提示:

1 <= numRows <= 30

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
一层层计算就行了，没啥特点
*/
func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	add := func(nums []int, i int) int {
		if i == 0 {
			return nums[0]
		}
		if i == len(nums) {
			return nums[len(nums)-1]
		}
		return nums[i-1] + nums[i]
	}
	res := [][]int{{1}}
	for i := 2; i <= numRows; i++ {
		temp := make([]int, i)
		for j := 0; j < i; j++ {
			temp[j] = add(res[i-2], j)
		}
		res = append(res, temp)
	}
	return res
}

/*问题*/
/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。



示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
每一个点处有两个选择，进或者不进，由此产生两个结果
迭代遍历，得到最终结果
问题在于过程中能否剪枝？
剪枝方案1：不可能连续三家都不偷
dp[i] = max(dp[i-2]+nums[i], dp[i-1])
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

/*问题*/
/*
超级丑数，给定一组质数，返回由这组质数作为质数因子构成的数组从小到大排列的第n个数
数组第一个数为1
*/
/*思路*/
/*
维护一个跟primers等长的数组pos，用于记录上一个由res[pos[i]] * primers[i]取得最小值的位置
*/
func superUgly(n int, primers []int) int {
	res := make([]int, n)
	res[0] = 1
	pos := make([]int, len(primers))
	for i := 1; i < n; i++ {
		curTermChooses := make([]int, 0)
		for ii, p := range pos {
			cur := res[p] * primers[ii]
			if res[i] == 0 {
				res[i] = cur
				curTermChooses = []int{ii}
				continue
			}

			if cur < res[i] {
				res[i] = cur
				curTermChooses = []int{ii}
				continue
			}

			// 这一段很关键，用于解决 2*3 = 3*2 的问题，也就是相同数字两种乘法都能同一步走到，这时候两种走法都应该+1
			if cur == res[i] {
				curTermChooses = append(curTermChooses, ii)
				continue
			}
		}
		for _, choose := range curTermChooses {
			pos[choose] = pos[choose] + 1
		}
	}

	return res[n-1]
}

/*问题*/
/*
王强今天很开心，公司发给N元的年终奖。王强决定把年终奖用于购物，他把想买的物品分为两类：主件与附件，附件是从属于某个主件的，下表就是一些主件与附件的例子：

主件	附件
电脑	打印机，扫描仪
书柜	图书
书桌	台灯，文具
工作椅	无

如果要买归类为附件的物品，必须先买该附件所属的主件。每个主件可以有 0 个、 1 个或 2 个附件。附件不再有从属于自己的附件。王强想买的东西很多，为了不超出预算，他把每件物品规定了一个重要度，分为 5 等：用整数 1 ~ 5 表示，第 5 等最重要。他还从因特网上查到了每件物品的价格（都是 10 元的整数倍）。他希望在不超过 N 元（可以等于 N 元）的前提下，使每件物品的价格与重要度的乘积的总和最大。
    设第 j 件物品的价格为 v[j] ，重要度为 w[j] ，共选中了 k 件物品，编号依次为 j 1 ， j 2 ，……， j k ，则所求的总和为：
v[j 1 ]*w[j 1 ]+v[j 2 ]*w[j 2 ]+ … +v[j k ]*w[j k ] 。（其中 * 为乘号）
    请你帮助王强设计一个满足要求的购物单。



输入描述：
输入的第 1 行，为两个正整数，用一个空格隔开：N m

（其中 N （ <32000 ）表示总钱数， m （ <60 ）为希望购买物品的个数。）


从第 2 行到第 m+1 行，第 j 行给出了编号为 j-1 的物品的基本数据，每行有 3 个非负整数 v p q


（其中 v 表示该物品的价格（ v<10000 ）， p 表示该物品的重要度（ 1 ~ 5 ）， q 表示该物品是主件还是附件。如果 q=0 ，表示该物品为主件，如果 q>0 ，表示该物品为附件， q 是所属主件的编号）




输出描述：
 输出文件只有一个正整数，为不超过总钱数的物品的价格与重要度乘积的总和的最大值（ <200000 ）。
示例1
输入：
1000 5
800 2 0
400 5 1
300 5 1
400 3 0
500 2 0
输出：
2200
*/
/*思路*/
/*
贪心算法可能有用吗？每次寻求最大价值
还是老实便利
剪支过程中，如果是从小到大排列，如果小的都买不了，大的也不用尝试了
附件怎么考虑，如果要买附件，就要考虑主键是否已经购买
也就是在做附件是否购买的抉择同时，买附件就意味着也买了主件，否则就不买附件
*/

type commondity struct {
	v        int
	p        int8
	selfInd  int
	ownerInd int
	isBuy    bool
}

func FindMaxValue(commondities []*commondity, leftMoney int) int {
	// // 价格从小到大排列
	// sort.Slice(commondities, func(i, j int) bool {
	// 	if commondities[i].v < commondities[j].v {
	// 		return true
	// 	}
	// 	if commondities[i].v > commondities[j].v {
	// 		return false
	// 	}
	// 	if commondities[i].ownerInd == 0 && commondities[j].ownerInd != 0 {
	// 		return true
	// 	}
	// 	if commondities[j].ownerInd == 0 && commondities[i].ownerInd != 0 {
	// 		return false
	// 	}
	// 	return true
	// })

	// 每个商品根据ind建立索引
	commondityMap := make(map[int]*commondity, len(commondities))
	for _, c := range commondities {
		commondityMap[c.selfInd] = c
	}

	return findMaxValue(commondities, leftMoney, 0, commondityMap)
}

func findMaxValue(commondities []*commondity, leftMoney int, cur int,
	commcommondityMap map[int]*commondity) int {
	if len(commondities) == 0 {
		return cur
	}
	if commondities[0].isBuy || commondities[0].v > leftMoney {
		return findMaxValue(commondities[1:], leftMoney, cur, commcommondityMap)
	}

	// 主件处理 或 附件对应主件已经购买
	if commondities[0].ownerInd == 0 || commcommondityMap[commondities[0].ownerInd].isBuy {
		// choose 1: 买第0个商品
		commcommondityMap[commondities[0].selfInd].isBuy = true
		choose1 := findMaxValue(commondities[1:], leftMoney-commondities[0].v,
			cur+commondities[0].v*int(commondities[0].p), commcommondityMap)
		commcommondityMap[commondities[0].selfInd].isBuy = false

		// choose 2: 不买第一个商品
		choose2 := findMaxValue(commondities[1:], leftMoney, cur, commcommondityMap)

		if choose1 >= choose2 {
			commcommondityMap[commondities[0].selfInd].isBuy = true
			return choose1
		}
		return choose2
	}

	// 附件处理 且 对应的主件还没买
	// 钱不够
	if commcommondityMap[commondities[0].ownerInd].v+commondities[0].v > leftMoney {
		return findMaxValue(commondities[1:], leftMoney, cur, commcommondityMap)
	}

	// 主附都买
	commcommondityMap[commondities[0].selfInd].isBuy = true
	commcommondityMap[commondities[0].ownerInd].isBuy = true
	choose1 := findMaxValue(commondities[1:], leftMoney-commondities[0].v-commcommondityMap[commondities[0].ownerInd].v,
		cur+commondities[0].v*int(commondities[0].p)+commcommondityMap[commondities[0].ownerInd].v*int(commcommondityMap[commondities[0].ownerInd].p),
		commcommondityMap)
	commcommondityMap[commondities[0].selfInd].isBuy = false
	commcommondityMap[commondities[0].ownerInd].isBuy = false

	// 附件不买
	choose2 := findMaxValue(commondities[1:], leftMoney, cur, commcommondityMap)
	if choose1 >= choose2 {
		commcommondityMap[commondities[0].selfInd].isBuy = true
		return choose1
	}
	return choose2
}

/*问题*/
/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。



示例 1：

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
示例 2：

输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。


提示：

1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ones-and-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
动态规划---背包问题
是否需要先排序？ 应该不需要
dp[i][cnt[0]][cnt[1]] = max(dp[i-1][cnt[0]][cnt[1]], dp[i][cnt[0]+getCnt0(i)][cnt[1]+getCnt1[i]])
*/
func findMaxForm(strs []string, m int, n int) int {
	getCnt := func(str string) (int, int) {
		cnt0, cnt1 := 0, 0
		for _, r := range str {
			if r == '1' {
				cnt1++
			} else {
				cnt0++
			}
		}
		return cnt0, cnt1
	}

	// 初始化
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for ii := range dp[i] {
			dp[i][ii] = make([]int, n+1)
		}
	}

	// 从第一个字符开始处理
	for i, str := range strs {
		cnt0, cnt1 := getCnt(str)
		for im := 0; im <= m; im++ {
			for in := 0; in <= n; in++ {
				dp[i+1][im][in] = dp[i][im][in]
				if im >= cnt0 && in >= cnt1 {
					tmp := dp[i][im-cnt0][in-cnt1] + 1
					if tmp > dp[i+1][im][in] {
						dp[i+1][im][in] = tmp
					}
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}
