package breadth_search

import (
	"math"
)

/*问题*/
/*
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。



示例 1：



输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：

输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
示例 3：

输入：[[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。


提示：

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] 仅为 0、1 或 2
*/
/*思路*/
/*
每分钟之前需要统计一次还有没有好苹果，是否需要遍历所有的？
不需要，已经走过的路径，其周围所有点肯定在下一次遍历中都走到了
两个map数组
a.目前为2的坐标
b.目前为1的坐标
每次循环目前为2的坐标，将其周围为1的坐标感染，并将被感染的坐标纳入第三个map
第三个mapc赋值给mapa，进行下一轮循环
循环中止条件：没有1了；循环一次后1的个数没变
*/
func OrangesRotting(grid [][]int) int {
	lengthRow := len(grid)
	lengthCol := len(grid[0])
	mapA := make(map[int]bool)
	mapB := make(map[int]bool)
	for i := 0; i < lengthRow; i++ {
		for j := 0; j < lengthCol; j++ {
			if grid[i][j] == 1 {
				mapB[i*100+j] = true
			}
			if grid[i][j] == 2 {
				mapA[i*100+j] = true
			}
		}
	}
	res := 0
	numNornal := len(mapB)
	for numNornal > 0 {
		tempNum := numNornal
		mapC := make(map[int]bool)
		for ind, _ := range mapA {
			if _, ok := mapC[ind]; !ok {
				row := ind / 100
				col := ind - row*100
				if _, ok := mapB[(row-1)*100+col]; ok {
					delete(mapB, (row-1)*100+col)
					tempNum--
					mapC[(row-1)*100+col] = true
				}
				if _, ok := mapB[(row+1)*100+col]; ok {
					delete(mapB, (row+1)*100+col)
					tempNum--
					mapC[(row+1)*100+col] = true
				}
				if _, ok := mapB[row*100+col-1]; ok {
					delete(mapB, row*100+col-1)
					tempNum--
					mapC[row*100+col-1] = true
				}
				if _, ok := mapB[row*100+col+1]; ok {
					delete(mapB, row*100+col+1)
					tempNum--
					mapC[row*100+col+1] = true
				}
			}
		}

		mapA = mapC

		if tempNum == numNornal {
			break
		}

		numNornal = tempNum
		res++
	}
	if numNornal > 0 {
		return -1
	}
	return res
}

/*问题*/
/*
让我们一起来玩扫雷游戏！

给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。

现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：

如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的方块都应该被递归地揭露。
如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
如果在此次点击中，若无更多方块可被揭露，则返回面板。


示例 1：

输入:

[['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'M', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E'],
 ['E', 'E', 'E', 'E', 'E']]

Click : [3,0]

输出:

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

解释:

示例 2：

输入:

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'M', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

Click : [1,2]

输出:

[['B', '1', 'E', '1', 'B'],
 ['B', '1', 'X', '1', 'B'],
 ['B', '1', '1', '1', 'B'],
 ['B', 'B', 'B', 'B', 'B']]

解释:



注意：

输入矩阵的宽和高的范围为 [1,50]。
点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
输入面板不会是游戏结束的状态（即有地雷已被挖出）。
简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。
*/
/*思路*/
/*
其实就是模拟点击了一次扫雷游戏之后，游戏内部的判断逻辑
采用广度优先搜索，从被点击的位置开始，递归上下左右全方位扫描
指定扫描顺序上下左右
1.如果周围都是'E'或者数字，则可以将该处设置成'B'，并且向上下左右延伸
2.如果周围有雷，更改该处为数字
3.如果某点处不是'E'或者'M'，说明该点已经搜索过，不需要再处理
*/

func updateBoard(board [][]byte, click []int) [][]byte {
	//判断是否中奖了
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	//没中，也就是'E'，遍历该点周围的点
	var checkRes byte
	var cnt byte = 0
	lPos := click[1] - 1
	rPos := click[1] + 1
	uPos := click[0] - 1
	dPos := click[0] + 1
	l := lPos >= 0
	r := rPos <= len(board[click[0]])-1
	u := uPos >= 0
	d := dPos <= len(board)-1
	//LU
	if l && u && board[uPos][lPos] == 'M' {
		cnt++
	}
	//U
	if u && board[uPos][click[1]] == 'M' {
		cnt++
	}
	//RU
	if r && u && board[uPos][rPos] == 'M' {
		cnt++
	}
	//L
	if l && board[click[0]][lPos] == 'M' {
		cnt++
	}
	//R
	if r && board[click[0]][rPos] == 'M' {
		cnt++
	}
	//LD
	if l && d && board[dPos][lPos] == 'M' {
		cnt++
	}
	//D
	if d && board[dPos][click[1]] == 'M' {
		cnt++
	}
	//RD
	if r && d && board[dPos][rPos] == 'M' {
		cnt++
	}
	if cnt == 0 {
		checkRes = 'B'
	} else {
		checkRes = cnt + '0'
	}

	board[click[0]][click[1]] = checkRes
	if checkRes != 'B' {
		return board
	}

	//碰到'B',再遍历click一圈
	//LU
	if l && u && (board[uPos][lPos] == 'M' || board[uPos][lPos] == 'E') {
		board = updateBoard(board, []int{uPos, lPos})
	}
	//U
	if u && (board[uPos][click[1]] == 'M' || board[uPos][click[1]] == 'E') {
		board = updateBoard(board, []int{uPos, click[1]})
	}
	//RU
	if r && u && (board[uPos][rPos] == 'M' || board[uPos][rPos] == 'E') {
		board = updateBoard(board, []int{uPos, rPos})
	}
	//L
	if l && (board[click[0]][lPos] == 'M' || board[click[0]][lPos] == 'E') {
		board = updateBoard(board, []int{click[0], lPos})
	}
	//R
	if r && (board[click[0]][rPos] == 'M' || board[click[0]][rPos] == 'E') {
		board = updateBoard(board, []int{click[0], rPos})
	}
	//LD
	if l && d && (board[dPos][lPos] == 'M' || board[dPos][lPos] == 'E') {
		board = updateBoard(board, []int{dPos, lPos})
	}
	//D
	if d && (board[dPos][click[1]] == 'M' || board[dPos][click[1]] == 'E') {
		board = updateBoard(board, []int{dPos, click[1]})
	}
	//RD
	if r && d && (board[dPos][rPos] == 'M' || board[dPos][rPos] == 'E') {
		board = updateBoard(board, []int{dPos, rPos})
	}
	return board
}

/*问题*/
/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1.暴力遍历，中途进行一些剪枝操作

2.动态规划
dp[i] = MIN(dp[i], dp[i - j * j] + 1)
*/
func NumSquaresWithCurNum(n int, curMin *int, curCnt int) {
	if *curMin != 0 && *curMin < curCnt {
		return
	}
	max := int(math.Floor(math.Sqrt(float64(n))))
	if max*max == n {
		if (*curMin == 0) || (*curMin > curCnt+1) {
			*curMin = curCnt + 1
		}
		return
	}
	for i := max; i >= 1; i-- {
		NumSquaresWithCurNum(n-i*i, curMin, curCnt+1)
	}
}

func NumSquaresV1(n int) int {
	curMin := new(int)
	NumSquaresWithCurNum(n, curMin, 0)
	return *curMin
}

func NumSquaresV2(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-j*j]+1)))
		}
	}
	return dp[n]
}
