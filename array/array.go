package array

import (
	"fmt"
	"github.com/bluesky1024/goAlgorithm/sort"
	"math"
	oriSort "sort"
	"strconv"
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

func (this *MountainArray) get(index int) int {
	return this.Arr[index]

}
func (this *MountainArray) length() int {
	return len(this.Arr)
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
		if a[0] == c[0] {
			return true
		}
		return false
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
	if num == 1 {
		return true
	}
	return false
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
	return
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
