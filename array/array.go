package array

import (
	"github.com/bluesky1024/goAlgorithm/sort"
	"math"
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
