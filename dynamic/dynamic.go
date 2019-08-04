package dynamic

import (
	"fmt"
	"math"
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

func (this *NumArray) SumRange(i int, j int) int {
	return this.SumInd[j] - this.SumInd[i] + this.Node[i]
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
