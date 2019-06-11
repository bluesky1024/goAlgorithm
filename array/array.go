package array

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
