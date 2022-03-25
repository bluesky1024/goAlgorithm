package array

import (
	"fmt"
	"math"
	"math/big"
	oriSort "sort"
	"strconv"

	"github.com/bluesky1024/goAlgorithm/sort"
)

/*问题*/
/*
最基础的问题：给定已排序数组，二分法查指定val位置，或者最近的一个数的位置
*/
/*思路*/
/*
没什么特别的，递归二分法
以下假设nums增序排列
*/
func FindNearestNum(nums []int, sInd int, eInd int, target int) (ind int) {
	if sInd == eInd {
		return sInd
	}
	if eInd-sInd == 1 {
		if (nums[eInd] - target) < (target - nums[sInd]) {
			return eInd
		} else {
			return sInd
		}
	}
	if nums[sInd] >= target {
		return sInd
	}
	if nums[eInd] <= target {
		return eInd
	}

	//查看中位数
	midInd := (sInd + eInd) / 2
	if nums[midInd] == target {
		return midInd
	}
	if nums[midInd] > target {
		return FindNearestNum(nums, sInd, midInd, target)
	} else {
		return FindNearestNum(nums, midInd, eInd, target)
	}
}

/*问题*/
/*
给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。

水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。

反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

示例 1:

输入: [[1,1,0],[1,0,1],[0,0,0]]
输出: [[1,0,0],[0,1,0],[1,1,1]]
解释: 首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
     然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
示例 2:

输入: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
输出: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
解释: 首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
     然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
说明:

1 <= A.length = A[0].length <= 20
0 <= A[i][j] <= 1
*/
/*思路*/
/*
1.水平翻转，用一个中间数组来替代
2.反转，直接取反 1-orig
3.返回值其实应该不是，但懒得改了，直接在A原值上修改
*/
func FlipAndInvertImage(A [][]int) [][]int {
	for ind, line := range A {
		length := len(line)
		temp := make([]int, length)
		for i := length - 1; i >= 0; i-- {
			temp[length-i-1] = 1 - line[i]
		}
		A[ind] = temp
	}
	return A
}

/*问题*/
/*
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。



示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。
*/
/*思路*/
/*
plan A:
一把梭，先平方，再快排
plan B:
前提：原数组已经排序
1.如果都小于0，逆序排列
2.如果都大于0，正序排列
3.如果跨0，分成两部分，相当于两个顺序链表的顺序合并
4.那么问题来了，怎么找到跨0的地方，考虑二分法的思路,0划分到正数的一方，分界线设定为离0最近的点
*/
func SortedSquares(A []int) []int {
	length := len(A)
	res := make([]int, length)
	partInd := FindNearestNum(A, 0, length-1, 0)

	le := 0

	if partInd == 0 {
		if A[partInd] >= 0 {
			for i, v := range A {
				res[i] = v * v
			}
			return res
		} else {
			le = partInd
		}
	} else if partInd == length-1 {
		if A[partInd] <= 0 {
			for i, v := range A {
				res[length-i-1] = v * v
			}
			return res
		} else {
			le = partInd - 1
		}
	} else {
		if A[partInd] >= 0 {
			le = partInd - 1
		} else {
			le = partInd
		}
	}

	left := A[:le+1]
	right := A[le+1:]
	lengthL := len(left)
	lengthR := len(right)
	resInd := 0

	i := lengthL - 1
	j := 0
	for {
		if i == -1 && j == lengthR {
			break
		}
		if i == -1 {
			res[resInd] = right[j] * right[j]
			resInd++
			j++
			continue
		}
		if j == lengthR {
			res[resInd] = left[i] * left[i]
			resInd++
			i--
			continue
		}

		if left[i]*left[i] <= right[j]*right[j] {
			res[resInd] = left[i] * left[i]
			i--
		} else {
			res[resInd] = right[j] * right[j]
			j++
		}
		resInd++
	}
	return res
}

/*问题*/
/*
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/
/*思路*/
/*
这道题没想过方案，是看了大佬的解题思路，简直精妙
从头到尾求累计乘积 然后反过来再求一次 每个位置的结果是两个方向的累计乘积乘起来
*/
func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	res[0] = 1
	for i := 1; i < length; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	temp := 1
	for i := length - 2; i >= 0; i-- {
		temp = temp * nums[i+1]
		res[i] = res[i] * temp
	}

	return res
}

/*问题*/
/*
学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。

请你返回至少有多少个学生没有站在正确位置数量。该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。



示例：

输入：[1,1,4,2,1,3]
输出：3
解释：
高度为 4、3 和最后一个 1 的学生，没有站在正确的位置。


提示：

1 <= heights.length <= 100
1 <= heights[i] <= 100
*/
/*思路*/
/*
非递减：似乎就是全部相邻数据满足相等或递增
先排序，排序数组跟原数组比较
不相等就表示位置不对
*/
func HeightChecker(heights []int) int {
	new_arr := append([]int{}, heights[:]...)
	sort.QuickSort(new_arr)
	length := len(new_arr)
	res := 0
	for i := 0; i < length; i++ {
		if new_arr[i] != heights[i] {
			res++
		}
	}
	return res
}

/*问题*/
/*
给出两个图像 A 和 B ，A 和 B 为大小相同的二维正方形矩阵。（并且为二进制矩阵，只包含0和1）。

我们转换其中一个图像，向左，右，上，或下滑动任何数量的单位，并把它放在另一个图像的上面。之后，该转换的重叠是指两个图像都具有 1 的位置的数目。

（请注意，转换不包括向任何方向旋转。）

最大可能的重叠是什么？

示例 1:

输入：A = [[1,1,0],
          [0,1,0],
          [0,1,0]]
     B = [[0,0,0],
          [0,1,1],
          [0,0,1]]
输出：3
解释: 将 A 向右移动一个单位，然后向下移动一个单位。
注意:

1 <= A.length = A[0].length = B.length = B[0].length <= 30
0 <= A[i][j], B[i][j] <= 1
*/
/*思路*/
/*
首先需要一个判断两个矩阵重合数的函数
其实应该考虑矩阵变换公式，但我忘了，目前因为没有旋转操作，可直接水平垂直遍历
*/
func getOverlap(A [][]int, B [][]int, lengthRow int, lengthCol int, i int, j int) int {
	res := 0
	for ii := 0; ii < lengthCol; ii++ {
		if ii+i < 0 {
			continue
		}
		if ii+i >= lengthCol {
			break
		}
		for jj := 0; jj < lengthRow; jj++ {
			if jj+j < 0 {
				continue
			}
			if jj+j >= lengthRow {
				break
			}
			if A[ii+i][jj+j] == 1 && B[ii][jj] == 1 {
				res++
			}
		}
	}
	return res
}

func LargestOverlap(A [][]int, B [][]int) int {
	lengthRow := len(A)
	lengthCol := len(A[0])
	max := 0
	for i := (-1*lengthRow + 1); i < lengthRow; i++ {
		for j := (-1*lengthCol + 1); j < lengthCol; j++ {
			temp := getOverlap(A, B, lengthRow, lengthCol, i, j)
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

/*问题*/
/*
 在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄，他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。现在，给出提莫对艾希的攻击时间序列和提莫攻击的中毒持续时间，你需要输出艾希的中毒状态总时长。

你可以认为提莫在给定的时间点进行攻击，并立即使艾希处于中毒状态。

示例1:

输入: [1,4], 2
输出: 4
原因: 在第 1 秒开始时，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒钟结束。
在第 4 秒开始时，提莫再次攻击艾希，使得艾希获得另外 2 秒的中毒时间。
所以最终输出 4 秒。
示例2:

输入: [1,2], 2
输出: 3
原因: 在第 1 秒开始时，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒钟结束。
但是在第 2 秒开始时，提莫再次攻击了已经处于中毒状态的艾希。
由于中毒状态不可叠加，提莫在第 2 秒开始时的这次攻击会在第 3 秒钟结束。
所以最终输出 3。
注意：

你可以假定时间序列数组的总长度不超过 10000。
你可以假定提莫攻击时间序列中的数字和提莫攻击的中毒持续时间都是非负整数，并且不超过 10,000,000。
*/
/*思路*/
/*
1.记录中毒总秒数
2.记录最近一次中毒有效期限
3.再被毒，视清空更新中毒事件
*/
func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0       //中毒秒数
	timeValid := 0 //到该秒时间结束，都是中毒状态
	for _, time := range timeSeries {
		if timeValid == 0 || time > timeValid {
			timeValid = time + duration - 1
			res = res + duration
			continue
		}
		res = res + time + duration - 1 - timeValid
		timeValid = time + duration - 1
	}
	return res
}

/*问题*/
/*
给你一个 山脉数组 mountainArr，请你返回能够使得 mountainArr.get(index) 等于 target 最小 的下标 index 值。

如果不存在这样的下标 index，就请返回 -1。



所谓山脉数组，即数组 A 假如是一个山脉数组的话，需要满足如下条件：

首先，A.length >= 3

其次，在 0 < i < A.length - 1 条件下，存在 i 使得：

A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]


你将 不能直接访问该山脉数组，必须通过 MountainArray 接口来获取数据：

MountainArray.get(k) - 会返回数组中索引为k 的元素（下标从 0 开始）
MountainArray.length() - 会返回该数组的长度


注意：

对 MountainArray.get 发起超过 100 次调用的提交将被视为错误答案。


示例 1：

输入：array = [1,2,3,4,5,3,1], target = 3
输出：2
解释：3 在数组中出现了两次，下标分别为 2 和 5，我们返回最小的下标 2。
示例 2：

输入：array = [0,1,2,4,2,1], target = 3
输出：-1
解释：3 在数组中没有出现，返回 -1。


提示：

3 <= mountain_arr.length() <= 10000
0 <= target <= 10^9
0 <= mountain_arr.get(index) <= 10^9
*/
/*思路*/
/*
第一反应，二分法
有个尖峰，所以数组由小到大，再到小，需要先找到最高峰
二分法无法判断最高峰再哪边，列举出各种情况
1,2,3,4,5,6,7,8,9,8,7,6      target=4
每一次取中位数，可能有三种可能
a.最高峰在 midInd 左边
b.midInd就是最高峰
c.最高峰在 midInd 右边

arr[midInd] <= arr[left] => a
arr[midInd] <= arr[right] => c
arr[midInd] > arr[left] && arr[midInd] > arr[right] => a,b,c
*/
type MountainArray struct {
	Arr []int
}

func (t *MountainArray) get(index int) int {
	return t.Arr[index]

}
func (t *MountainArray) length() int {
	return len(t.Arr)
}

func FindMoutain(mountainArray *MountainArray, left int, right int, leftV int, rightV int) (int, int) {
	if right-left <= 1 {
		if leftV > rightV {
			return left, leftV
		} else {
			return right, rightV
		}
	}
	midInd := int(math.Floor(float64(left+right) / 2))

	midV := mountainArray.get(midInd)
	if midV <= leftV {
		return FindMoutain(mountainArray, left, midInd, leftV, midV)
	}
	if midV <= rightV {
		return FindMoutain(mountainArray, midInd, right, midV, rightV)
	}
	if midV < mountainArray.get(midInd+1) {
		return FindMoutain(mountainArray, midInd, right, midV, rightV)
	} else {
		return FindMoutain(mountainArray, left, midInd, leftV, midV)
	}
}

func FindTargetInSort(order bool, mountainArray *MountainArray, left int, right int, leftV int, rightV int, target int) int {
	if right-left <= 1 {
		if rightV == target {
			return right
		}
		if leftV == target {
			return left
		}
		return -1
	}
	midInd := int(math.Floor(float64(left+right) / 2))

	midV := mountainArray.get(midInd)
	if midV == target {
		return midInd
	}
	if order {
		if midV < target {
			return FindTargetInSort(order, mountainArray, midInd, right, midV, rightV, target)
		}

		return FindTargetInSort(order, mountainArray, left, midInd, leftV, midV, target)
	}

	if midV > target {
		return FindTargetInSort(order, mountainArray, midInd, right, midV, rightV, target)
	}

	return FindTargetInSort(order, mountainArray, left, midInd, leftV, midV, target)
}

func FindInMountainArray(target int, mountainArr *MountainArray) int {
	len := mountainArr.length()
	left := 0
	right := len - 1
	leftV := mountainArr.get(left)
	rightV := mountainArr.get(right)
	midInd, midV := FindMoutain(mountainArr, left, right, leftV, rightV)
	res := FindTargetInSort(true, mountainArr, left, midInd, leftV, midV, target)
	if res != -1 {
		return res
	}
	return FindTargetInSort(false, mountainArr, midInd, right, midV, rightV, target)
}

/*问题*/
/*
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

示例 1:

输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4
示例 2:

输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
*/
/*思路*/
/*
1.需要一个方法判断三个点共线，判断过程因为不能用分数，改成交叉相乘判断是否相等来判断共线
2.遍历任意两点构成的直线覆盖了多少个点，统计max
*/
//判断是否三点共线
func checkInLine(a []int, b []int, c []int) bool {
	if a[0] == b[0] {
		return a[0] == c[0]
	}
	if a[0] == c[0] {
		return false
	}

	if (a[0]-b[0])*(a[1]-c[1]) == (a[1]-b[1])*(a[0]-c[0]) {
		return true
	}
	return false
}

//获取aInd,bInd构成的直线上的点个数
func getPointInLineNum(aInd int, bInd int, points [][]int) int {
	res := 2
	for ind, point := range points {
		if ind == aInd || ind == bInd {
			continue
		}
		if (point[0] == points[aInd][0] && point[1] == points[aInd][1]) || (point[0] == points[bInd][0] && point[1] == points[bInd][1]) {
			res++
			continue
		}
		if checkInLine(points[aInd], points[bInd], point) {
			res++
		}
	}
	return res
}

//判断是否重复直线
func checkRepeatLine(lines [][]int, aInd int, bInd int, points [][]int) bool {
	for _, line := range lines {
		if checkInLine(points[line[0]], points[line[1]], points[aInd]) && checkInLine(points[line[0]], points[line[1]], points[bInd]) {
			return true
		}
	}
	return false
}

func MaxPoints(points [][]int) int {
	pointNum := len(points)
	if pointNum <= 2 {
		return pointNum
	}

	////存储已经遍历过的直线
	//lines := make([][]int, 0)

	max := 0
	for i := 0; i < pointNum; i++ {
		for j := i + 1; j < pointNum; j++ {
			//if checkRepeatLine(lines, i, j, points) {
			//	continue
			//}
			temp := getPointInLineNum(i, j, points)
			if max < temp {
				max = temp
			}
			//lines = append(lines, []int{i, j})
		}
	}
	return max
}

/*问题*/
/*
编写一个程序判断给定的数是否为丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例 1:

输入: 6
输出: true
解释: 6 = 2 × 3
示例 2:

输入: 8
输出: true
解释: 8 = 2 × 2 × 2
示例 3:

输入: 14
输出: false
解释: 14 不是丑数，因为它包含了另外一个质因数 7。
说明：

1 是丑数。
输入不会超过 32 位有符号整数的范围: [−231,  231 − 1]。
*/
/*思路*/
/*
循环除2，3，5就完事了
*/
func IsUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for num%2 == 0 {
		num = num / 2
	}
	for num%3 == 0 {
		num = num / 3
	}
	for num%5 == 0 {
		num = num / 5
	}
	return num == 1
}

/*问题*/
/*
根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在1970年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞具有一个初始状态 live（1）即为活细胞， 或 dead（0）即为死细胞。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
根据当前状态，写一个函数来计算面板上细胞的下一个（一次更新后的）状态。下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。

示例:

输入:
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
输出:
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
进阶:

你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/game-of-life
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*

 */
func checkLife(row int, col int, lengthRow int, lengthCol int, board [][]int) int {
	cnt := 0
	l := false
	r := false
	u := false
	d := false
	if col != 0 {
		l = true
	}
	if col+1 < lengthCol {
		r = true
	}
	if row != 0 {
		u = true
	}
	if row+1 < lengthRow {
		d = true
	}

	if l && board[row][col-1] == 1 {
		cnt++
	}
	if r && board[row][col+1] == 1 {
		cnt++
	}
	if u && board[row-1][col] == 1 {
		cnt++
	}
	if d && board[row+1][col] == 1 {
		cnt++
	}
	if l && u && board[row-1][col-1] == 1 {
		cnt++
	}
	if l && d && board[row+1][col-1] == 1 {
		cnt++
	}
	if r && u && board[row-1][col+1] == 1 {
		cnt++
	}
	if r && d && board[row+1][col+1] == 1 {
		cnt++
	}

	return cnt
}

func GameOfLifeV1(board [][]int) {
	lengthRow := len(board)
	if lengthRow == 0 {
		return
	}
	lengthCol := len(board[0])
	if lengthCol == 0 {
		return
	}

	newBoard := make([][]int, lengthRow)
	for row, line := range board {
		temp := make([]int, lengthCol)
		for col, v := range line {
			cnt := checkLife(row, col, lengthRow, lengthCol, board)
			if v == 1 {
				if cnt > 3 || cnt < 2 {
					temp[col] = 0
					continue
				}
				temp[col] = 1
			} else {
				if cnt == 3 {
					temp[col] = 1
					continue
				}
				temp[col] = 0
			}
		}
		newBoard[row] = temp
	}

	for i := 0; i < lengthRow; i++ {
		for j := 0; j < lengthCol; j++ {
			board[i][j] = newBoard[i][j]
		}
	}
}

/*问题*/
/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1.三个for循环暴力遍历
2.加个map，可以少一层循环
*/
func ThreeSumV1(nums []int) [][]int {
	oriSort.Ints(nums)
	length := len(nums)
	res := make([][]int, 0)
	checkRes := make(map[string]bool)
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < length-1; j++ {
			if nums[i]+nums[j]*2 > 0 {
				break
			}
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] > 0 {
					break
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					tempStr := strconv.Itoa(nums[i]) + "_" + strconv.Itoa(nums[j]) + "_" + strconv.Itoa(nums[k])
					if _, ok := checkRes[tempStr]; !ok {
						checkRes[tempStr] = true
						res = append(res, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return res
}

func ThreeSumV2(nums []int) [][]int {
	oriSort.Ints(nums)
	fmt.Println(nums)
	length := len(nums)
	res := make([][]int, 0)
	checkRes := make(map[string]bool)
	checkNum := make(map[int]int)
	for _, v := range nums {
		if cnt, ok := checkNum[v]; ok {
			checkNum[v] = cnt + 1
		} else {
			checkNum[v] = 1
		}
	}
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < length; j++ {
			if (nums[i] + nums[j]*2) > 0 {
				break
			}
			thirdOne := 0 - nums[i] - nums[j]
			if cnt, ok := checkNum[thirdOne]; ok {
				if thirdOne == nums[j] && thirdOne != nums[i] && cnt < 2 {
					break
				}
				if thirdOne == nums[j] && thirdOne == nums[i] && cnt < 3 {
					break
				}
				tempStr := strconv.Itoa(nums[i]) + "_" + strconv.Itoa(nums[j]) + "_" + strconv.Itoa(thirdOne)
				if _, ok := checkRes[tempStr]; !ok {
					checkRes[tempStr] = true
					res = append(res, []int{nums[i], nums[j], thirdOne})
				}
			}
		}
	}
	return res
}

/*问题*/
/*
给定一个整数 n, 返回从 1 到 n 的字典顺序。

例如，

给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。

请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lexicographical-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
字典序是确定的，可以从1开始遍历
遍历顺序无非就是每次*10+i，i从1到9
终止条件遍历数大于n
只是这里面的res是不停的在进行append和copy，内存占用其实很高
要进行优化，可以定义一个新的结构体包裹切片，这个结构体变量的指针作为传入参数，所有迭代过程对同一个变量进行操作，可以减少copy操作，同时内存占用减少
*/
func LexicalOrderWithBase(n int, base int) []int {
	res := make([]int, 0)
	base = base * 10
	var i int
	if base == 0 {
		i = 1
	} else {
		i = 0
	}
	for ; i < 10; i++ {
		if base+i > n {
			break
		}
		res = append(res, base+i)
		res = append(res, LexicalOrderWithBase(n, base+i)...)
	}
	return res
}
func LexicalOrder(n int) []int {
	res := LexicalOrderWithBase(n, 0)
	return res
}

/*问题*/
/*
根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
直观来看，第i位的数据要与后续n-i位有关
如果反过来进行遍历，可以减少遍历次数，第i位的数据首先参考第i+1位，如果不对，跳转到新数组中第i+1位的数据
*/
func DailyTemperatures(T []int) []int {
	length := len(T)
	if length == 0 {
		return nil
	}
	res := make([]int, length)
	res[length-1] = 0
	for i := length - 2; i >= 0; i-- {
		checkInd := i + 1
		temp := T[checkInd]
		for {
			if temp > T[i] {
				res[i] = checkInd - i
				break
			}
			if res[checkInd] == 0 {
				res[i] = 0
				break
			}
			checkInd = res[checkInd] + checkInd
			temp = T[checkInd]
		}
	}
	return res
}

type SumRes struct {
	arr [][]int
}

func getResWithIn(nums []int, curInd int, used []int, t int, res *SumRes) {
	if t == 0 {
		res.arr = append(res.arr, used)
	}
	length := len(nums)
	for i := curInd; i < length; i++ {
		if nums[i] > t {
			break
		}
		temp := append(used, nums[i])
		getResWithIn(nums, i+1, temp, t-nums[i], res)
	}
}

func GetRes(nums []int, t int) [][]int {
	oriSort.Ints(nums)

	res := &SumRes{
		arr: make([][]int, 0),
	}
	used := make([]int, 0)
	getResWithIn(nums, 0, used, t, res)

	return res.arr
}

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

func DecimalToAny(num int, n int) string {
	if n < 2 || n > 76 {
		return ""
	}

	res := ""
	for num != 0 {
		temp := num % n
		tempStr := tenToAny[temp]

		res = tempStr + res
		fmt.Println("temp:", temp, "num:", num)
		num = num / n
	}
	return res
}

/*问题*/
/*
给你一个整数 n 和一个整数数组 rounds 。有一条圆形赛道由 n 个扇区组成，扇区编号从 1 到 n 。现将在这条赛道上举办一场马拉松比赛，该马拉松全程由 m 个阶段组成。其中，第 i 个阶段将会从扇区 rounds[i - 1] 开始，到扇区 rounds[i] 结束。举例来说，第 1 阶段从 rounds[0] 开始，到 rounds[1] 结束。

请你以数组形式返回经过次数最多的那几个扇区，按扇区编号 升序 排列。

注意，赛道按扇区编号升序逆时针形成一个圆（请参见第一个示例）。



示例 1：



输入：n = 4, rounds = [1,3,1,2]
输出：[1,2]
解释：本场马拉松比赛从扇区 1 开始。经过各个扇区的次序如下所示：
1 --> 2 --> 3（阶段 1 结束）--> 4 --> 1（阶段 2 结束）--> 2（阶段 3 结束，即本场马拉松结束）
其中，扇区 1 和 2 都经过了两次，它们是经过次数最多的两个扇区。扇区 3 和 4 都只经过了一次。
示例 2：

输入：n = 2, rounds = [2,1,2,1,2,1,2,1,2]
输出：[2]
示例 3：

输入：n = 7, rounds = [1,3,5,7]
输出：[1,2,3,4,5,6,7]


提示：

2 <= n <= 100
1 <= m <= 100
rounds.length == m + 1
1 <= rounds[i] <= n
rounds[i] != rounds[i + 1] ，其中 0 <= i < m
*/
/*思路*/
/*
这道题最难的是理解题意

*/
func MostVisited(n int, rounds []int) []int {
	length := len(rounds)
	cntMap := make(map[int]int)
	for i := 1; i <= n; i++ {
		cntMap[i] = 0
	}
	for i := 1; i < length; i++ {
		s := rounds[i-1]
		e := rounds[i]

		//s -> e左闭右开区间每个数都+1
		if s < e {
			for ind := s; ind < e; ind++ {
				cntMap[ind] = cntMap[ind] + 1
			}
		}

		//e -> n -> 1->s 左闭右开区间每个数都+1
		if s > e {
			for ind := s; ind <= n; ind++ {
				cntMap[ind] = cntMap[ind] + 1
			}
			for ind := 1; i < e; ind++ {
				cntMap[ind] = cntMap[ind] + 1
			}
		}
	}

	//右开区间，所以补上最后一次的计数
	cntMap[rounds[length-1]] = cntMap[rounds[length-1]] + 1

	//找出最大值
	curMax := 0
	reverCntMap := make(map[int][]int)
	for ind, cnt := range cntMap {
		reverCntMap[cnt] = append(reverCntMap[cnt], ind)
		if curMax < cnt {
			curMax = cnt
		}
	}

	return reverCntMap[curMax]
}

/*问题*/
/*
在一排座位（ seats）中，1 代表有人坐在座位上，0 代表座位上是空的。

至少有一个空座位，且至少有一人坐在座位上。

亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。

返回他到离他最近的人的最大距离。



示例 1：

输入：[1,0,0,0,1,0,1]
输出：2
解释：
如果亚历克斯坐在第二个空位（seats[2]）上，他到离他最近的人的距离为 2 。
如果亚历克斯坐在其它任何一个空位上，他到离他最近的人的距离为 1 。
因此，他到离他最近的人的最大距离是 2 。
示例 2：

输入：[1,0,0,0]
输出：3
解释：
如果亚历克斯坐在最后一个座位上，他离最近的人有 3 个座位远。
这是可能的最大距离，所以答案是 3 。


提示：

2 <= seats.length <= 20000
seats 中只含有 0 和 1，至少有一个 0，且至少有一个 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximize-distance-to-closest-person
*/
/*思路*/
/*
正向遍历一次，再反向遍历一次，取两者中较小值o(n)
    1,0,0,0,1,0,1
->  0 1 2 3 0 1 0
    0 3 2 1 0 1 0  <-

    0 1 2 1 0 1 0
max 2

有个细节注意点：如果第一位就是0，则按顺序遍历可能出现距离无穷大的情况，在代码实现中用-1表示无穷大
*/
func MaxDistToClosest(seats []int) int {
	disList := make([]int, len(seats))
	//正向遍历
	curDis := -1 //用-1表示无穷大
	for i, d := range seats {
		if d == 1 {
			curDis = 0
			disList[i] = 0
			continue
		}
		if curDis >= 0 {
			curDis++
		}
		disList[i] = curDis
	}

	//反向遍历
	curDis = -1
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			curDis = 0
			disList[i] = 0
			continue
		}
		if curDis >= 0 {
			curDis++

			if curDis < disList[i] || disList[i] == -1 {
				disList[i] = curDis
			}
		}
		//else {
		//	//curDis还是-1无穷大，则距离一定是用正向的距离，不需要比较
		//	disList[]
		//}
	}

	//遍历disList,取得最大值
	max := 0
	for _, d := range disList {
		if d > max {
			max = d
		}
	}

	return max
}

/*题目*/
/*
在二维数组grid中，grid[i][j]代表位于某处的建筑物的高度。 我们被允许增加任何数量（不同建筑物的数量可能不同）的建筑物的高度。 高度 0 也被认为是建筑物。

最后，从新数组的所有四个方向（即顶部，底部，左侧和右侧）观看的“天际线”必须与原始数组的天际线相同。 城市的天际线是从远处观看时，由所有建筑物形成的矩形的外部轮廓。 请看下面的例子。

建筑物高度可以增加的最大总和是多少？

例子：
输入： grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
输出： 35
解释：
The grid is:
[ [3, 0, 8, 4],
  [2, 4, 5, 7],
  [9, 2, 6, 3],
  [0, 3, 1, 0] ]

从数组竖直方向（即顶部，底部）看“天际线”是：[9, 4, 8, 7]
从水平水平方向（即左侧，右侧）看“天际线”是：[8, 7, 9, 3]

在不影响天际线的情况下对建筑物进行增高后，新数组如下：

gridNew = [ [8, 4, 8, 7],
            [7, 4, 7, 7],
            [9, 4, 8, 7],
            [3, 3, 3, 3] ]
说明:

1 < grid.length = grid[0].length <= 50。
 grid[i][j] 的高度范围是： [0, 100]。
一座建筑物占据一个grid[i][j]：换言之，它们是 1 x 1 x grid[i][j] 的长方体。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
算出横竖各自最大值
每个楼层的提升空间受x轴、y轴两方面制约，取其中较小的
*/

func MaxIncreaseKeepingSkyline(grid [][]int) int {
	findMin := func(a int, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	udMax := make([]int, len(grid[0]))
	lrMax := make([]int, len(grid))
	// 找到横看竖看各自的最大值列表
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > udMax[j] {
				udMax[j] = grid[i][j]
			}
			if grid[i][j] > lrMax[i] {
				lrMax[i] = grid[i][j]
			}
		}
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			curLimit := findMin(udMax[j], lrMax[i])
			if grid[i][j] < curLimit {
				res += curLimit - grid[i][j]
			}
		}
	}

	return res
}

/*问题*/
/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
字典序可以理解为从后往前算
通过递减冒泡排序，每多一个数进行一次插入排序，并与原数组比较，若不同，则可直接输出
若遍历到最后都相同，则数组翻转
*/

func nextPermutationX(nums []int) {
	if len(nums) <= 1 {
		return
	}

	hasChange := false
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i; j < len(nums)-1; j++ {
			if nums[j] < nums[j+1] {
				hasChange = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		// 若调整过
		if hasChange {
			return
		}
	}

	// 遍历结束，说明没有调整
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	hasChange := false
	for i := len(nums) - 2; i >= 0; i-- {
		var j int
		for j = len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				hasChange = true
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
		if hasChange {
			oriSort.Ints(nums[i+1:])
			return
		}
	}

	// 遍历结束，说明没有调整
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

/*问题*/
/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
什么情况下会需要合并？
a区间的enda 大于b区间的startb
合并后的区间为 min(starta,startb) - max(enda,endb)
*/

type itv struct {
	data [][]int
}

func (i *itv) Len() int {
	return len(i.data)
}
func (i *itv) Less(ii, jj int) bool {
	return i.data[ii][0] < i.data[jj][0]
}
func (i *itv) Swap(ii, jj int) {
	i.data[ii], i.data[jj] = i.data[jj], i.data[ii]
}

func merge(intervals [][]int) [][]int {
	// 对 intervals 进行排序
	itvData := &itv{
		data: intervals,
	}

	oriSort.Sort(itvData)

	res := make([][]int, 0)
	for _, interval := range intervals {
		res = mergeSingle(res, interval)
	}

	return res
}

func mergeSingle(resIntervals [][]int, interval []int) [][]int {
	temp := interval
	getMin := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	getMax := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := make([][]int, 0)
	insert := false
	curInd := 0
	for {
		if curInd >= len(resIntervals) {
			break
		}
		for ind := curInd; ind < len(resIntervals); ind++ {
			itv := resIntervals[ind]
			// 有重合
			if interval[0] >= itv[0] && interval[0] <= itv[1] {
				res = append(res, []int{getMin(itv[0], interval[0]), getMax(itv[1], interval[1])})
				curInd = ind
				interval = res[len(res)-1]
				insert = true
				break
			} else {
				res = append(res, itv)
			}
			curInd = ind
		}
		curInd++
	}

	if !insert {
		res = append(res, temp)
	}

	return res
}

/*问题*/
/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。



示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
逐个判断每个节点最多能到哪个点
从前往后遍历，最后判断
*/

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	curCan := make([]bool, len(nums))
	curCan[0] = true
	for i := range nums {
		if !curCan[i] {
			return false
		}
		for ind := i + 1; ind <= i+nums[i] && ind <= len(nums)-1; ind++ {
			curCan[ind] = true
		}
		if curCan[len(nums)-1] {
			return true
		}
	}

	return curCan[len(nums)-1]
}

/*问题*/
/*
峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。



示例 1：

输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例 2：

输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-peak-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
主要难点是要logn的复杂度
参考二分法，通过逐步缩减查询范围，来进行查询
相对于从前往后遍历，更为稳定
*/
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	isLeftPeak := func(numsInner []int, i int) bool {
		if (i == 0) || (numsInner[i] > numsInner[i-1]) {
			return true
		}
		return false
	}
	isRightPeak := func(numsInner []int, i int) bool {
		if (i == len(numsInner)-1) || (numsInner[i] > numsInner[i+1]) {
			return true
		}
		return false
	}

	head := 0
	tail := len(nums) - 1
	for {
		mid := (head + tail) / 2
		if !isLeftPeak(nums, mid) {
			tail = mid - 1
			continue
		}
		if !isRightPeak(nums, mid) {
			head = mid + 1
			continue
		}
		return mid
	}
}

/*问题*/
/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。



示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：

输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 3：

输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
示例 4：

输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。


提示：

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
nums 中的所有数字都 独一无二

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
根据数组长度，确定 n
怎么判断哪个数不存在
总和是 (1+n)*n/2 + missNum
length == 1 的时候，居然也成立，那就不用特殊处理了
*/
func missingNumber(nums []int) int {
	length := len(nums)
	res := (1 + length) * length / 2
	for _, num := range nums {
		res -= num
	}
	return res
}

/*问题*/
/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。



示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0


提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1. 排序
2. 固定第一个数，其他数从该数后面选
3. 双指针从头和尾开始遍历，从外往内遍历进行剪支
4. 需要注意：以1 2 3 4 5 6 7 8 9 为例
   2+8 可能大于 3+4 也可能 小于 5+6
   也就是，不能直接在发现数据差绝对值变大之后就停止遍历
*/
func threeSumClosest(nums []int, target int) int {
	// 排序
	oriSort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	// 遍历第一个数字
	for i := 0; i < len(nums)-2; i++ {
		head := i + 1
		tail := len(nums) - 1
		curSum := 0
		for head != tail {
			curSum = nums[i] + nums[head] + nums[tail]
			if math.Abs(float64(target-curSum))-math.Abs(float64(target-res)) < 0 {
				res = curSum
			}
			// 如果相等，则直接命中target
			if nums[head]+nums[tail] == target-nums[i] {
				return target
			}
			// 如果小于 left,说明数小了
			if nums[head]+nums[tail] < target-nums[i] {
				head++
				continue
			}
			// 如果大于 left,说明数大了
			if nums[head]+nums[tail] > target-nums[i] {
				tail--
				continue
			}
		}
	}
	return res
}

/*问题*/
/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4


进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
7 0 1 2 4 5 6             0


4 5 6 7 0 1 2             0

4 2 7
mid > tail head=>mid 是升序
mid < tail mid=>tail 是升序
*/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

/*问题*/
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？



示例 1：


输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6


提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1. 递归
2. 排列组合
*/
func uniquePathsV1(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePathsV1(m-1, n) + uniquePathsV1(m, n-1)
}

func uniquePathsV2(m int, n int) int {
	// C(m+n-2,m-1)
	// (m+n-2)!/((m-1)! * (n-1)!)
	return int(big.NewInt(1).Binomial(int64(m+n-2), int64(m-1)).Int64())
}

/*问题*/
/*
给定一组正整数，相邻的整数之间将会进行浮点除法操作。例如， [2,3,4] -> 2 / 3 / 4 。

但是，你可以在任意位置添加任意数目的括号，来改变算数的优先级。你需要找出怎么添加括号，才能得到最大的结果，并且返回相应的字符串格式的表达式。你的表达式不应该含有冗余的括号。

示例：

输入: [1000,100,10,2]
输出: "1000/(100/10/2)"
解释:
1000/(100/10/2) = 1000/((100/10)/2) = 200
但是，以下加粗的括号 "1000/((100/10)/2)" 是冗余的，
因为他们并不影响操作的优先级，所以你需要返回 "1000/(100/10/2)"。

其他用例:
1000/(100/10)/2 = 50
1000/(100/(10/2)) = 50
1000/100/10/2 = 0.5
1000/100/(10/2) = 2
说明:

输入数组的长度在 [1, 10] 之间。
数组中每个元素的大小都在 [2, 1000] 之间。
每个测试用例只有一个最优除法解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/optimal-division
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
nums 中 第一个数一定是分子，第二个数一定是分母
其余数据可以在分子和分母之间任意切换，其他数全都是分子即可
*/
func optimalDivision(nums []int) string {
	res := ""
	for i, num := range nums {
		if i == 0 {
			res += fmt.Sprintf("%d/(", num)
			continue
		}
		if i == len(nums)-1 {
			res += fmt.Sprintf("%d)", num)
			continue
		}
		res += fmt.Sprintf("%d/", num)
	}
	return res
}

/*
全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getLeftNums(nums []int, deleteNum int) []int {
	res := make([]int, 0, len(nums)-1)
	for _, num := range nums {
		if num != deleteNum {
			res = append(res, num)
		}
	}
	return res
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		tmpRes := permute(getLeftNums(nums, nums[i]))
		for _, tmpNums := range tmpRes {
			res = append(res, append([]int{nums[i]}, tmpNums...))
		}
	}
	return res
}

/*问题*/
/*
全排列 2.0
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getLeftNumsUnique(nums []int, deleteNum int) []int {
	res := make([]int, 0, len(nums)-1)
	isFilter := false
	for _, num := range nums {
		if !isFilter && num == deleteNum {
			isFilter = true
			continue
		}
		res = append(res, num)
	}
	return res
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	res := make([][]int, 0)
	checkFirst := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := checkFirst[nums[i]]; ok {
			continue
		}
		checkFirst[nums[i]] = struct{}{}
		tmpRes := permuteUnique(getLeftNumsUnique(nums, nums[i]))
		for _, tmpNums := range tmpRes {
			res = append(res, append([]int{nums[i]}, tmpNums...))
		}
	}
	return res
}
