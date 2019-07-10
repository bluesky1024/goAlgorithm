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
你的任务是计算 a^b 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。

示例 1:

输入: a = 2, b = [3]
输出: 8
示例 2:

输入: a = 2, b = [1,0]
输出: 1024
*/
/*思路*/
/*
 */
