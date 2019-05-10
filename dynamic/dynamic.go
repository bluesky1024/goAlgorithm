package dynamic

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
	Node []int
}

func ConstructorNumArray(nums []int) NumArray {
	tempSum := 0
	res := NumArray{}
	res.SumInd = make([]int,len(nums))
	res.Node = make([]int,len(nums))
	for ind,v:=range nums{
		res.Node[ind] = v
		tempSum = tempSum + v
		res.SumInd[ind] = tempSum
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.SumInd[j]-this.SumInd[i] + this.Node[i]
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
func getMinLastStep(minSumArr [][]int,row int,col int) int{
	if row == 0 {
		if col == 0{
			return 0
		}
		return minSumArr[row][col-1]
	}
	if col == 0 {
		return minSumArr[row-1][col]
	}

	if minSumArr[row-1][col] <= minSumArr[row][col-1]{
		return minSumArr[row-1][col]
	}
	return minSumArr[row][col-1]
}

func MinPathSum(grid [][]int) int {
	minSumArr := make([][]int,len(grid))
	for rowInd,arr := range grid{
		minSumArr[rowInd] = make([]int,len(arr))
		for colInd,v := range arr {
			minSumArr[rowInd][colInd] = getMinLastStep(minSumArr,rowInd,colInd) + v
		}
	}
	return minSumArr[len(minSumArr)-1][len(minSumArr[0])-1]
}
