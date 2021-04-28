package breadth_search

import (
	"fmt"
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

/*问题*/
/*
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
有点像围棋？只是边界上不算被包围
所以只要某块联通的白棋能延伸到边界上就算自由棋
广度遍历，上下左右一路找下去，找到一团互联的白棋

所以问题是遍历顺序，怎么记录互联关系，怎么记录哪些棋子已经被遍历过了

然后发现思维僵化，应该反向寻找，在边界上找白棋，从边界白棋开始向内扩展，凡是扩展到的白棋，都是自由棋，合理

存储的话，可以将自由棋先换成 非 'X' 和 'O' 的数值，如'N'
待四个边界遍历完毕，再按层序遍历矩阵将所有的非'N'换成'X',将所有的'N'换成'O'
*/
func XOChessSolve(board [][]byte) {
	length := len(board)
	if length == 0 {
		return
	}

	width := len(board[0])
	//上边界
	for i := 0; i < length; i++ {
		if board[i][0] == 'O' {
			XOChessSearch(board, length, width, i, 0)
		}
	}
	fmt.Println("----------")
	for i := 0; i < length; i++ {
		fmt.Println(board[i])
	}
	//下边界
	for i := 0; i < length; i++ {
		if board[i][width-1] == 'O' {
			XOChessSearch(board, length, width, i, width-1)
		}
	}
	fmt.Println("----------")
	for i := 0; i < length; i++ {
		fmt.Println(board[i])
	}
	//左边界
	for i := 0; i < width; i++ {
		if board[0][i] == 'O' {
			XOChessSearch(board, length, width, 0, i)
		}
	}
	fmt.Println("----------")
	for i := 0; i < length; i++ {
		fmt.Println(board[i])
	}
	//右边界
	for i := 0; i < width; i++ {
		if board[length-1][i] == 'O' {
			XOChessSearch(board, length, width, length-1, i)
		}
	}

	fmt.Println("----------")
	for i := 0; i < length; i++ {
		fmt.Println(board[i])
	}

	//数据转换
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 'N' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func XOChessSearch(board [][]byte, length int, width int, curI int, curJ int) {
	if board[curI][curJ] == 'X' || board[curI][curJ] == 'N' {
		return
	}
	board[curI][curJ] = 'N'
	if curI-1 >= 0 {
		XOChessSearch(board, length, width, curI-1, curJ)
	}
	if curJ-1 >= 0 {
		XOChessSearch(board, length, width, curI, curJ-1)
	}
	if curI+1 < length {
		XOChessSearch(board, length, width, curI+1, curJ)
	}
	if curJ+1 < width {
		XOChessSearch(board, length, width, curI, curJ+1)
	}
}

/*问题*/
/*
用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。



示例 1：



输入：n = 4, connections = [[0,1],[0,2],[1,2]]
输出：1
解释：拔下计算机 1 和 2 之间的线缆，并将它插到计算机 1 和 3 上。
示例 2：



输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
输出：2
示例 3：

输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
输出：-1
解释：线缆数量不足。
示例 4：

输入：n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
输出：0


提示：

1 <= n <= 10^5
1 <= connections.length <= min(n*(n-1)/2, 10^5)
connections[i].length == 2
0 <= connections[i][0], connections[i][1] < n
connections[i][0] != connections[i][1]
没有重复的连接。
两台计算机不会通过多条线缆连接。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected
*/
/*思路*/
/*
存下来目前条件下可达的节点集，如果某条线两端的节点都已经在某个可达节点集里，则这条边就可以删除，也就是可以给别人用

遍历所有的边，看看最终组合成多少个互相独立的节点集，单个节点不跟外界沟通的独立成为一个节点集
假设总共n个独立的集合的话，需要n-1条数据连接
然后判断上述删除的多余边是否>=n-1,若是则返回n-1，若否，则返回-1
*/
func findNumInSet(num int, setList [][]int) int {
	for ind, set := range setList {
		for _, n := range set {
			if num == n {
				return ind
			}
		}
	}

	return -1
}

func MakeConnected(n int, connections [][]int) int {
	setList := make([][]int, 0)

	//遍历找到所有的互联小集合,同时找出多余线
	invalidConnCnt := 0
	for _, connection := range connections {
		p1 := findNumInSet(connection[0], setList)
		p2 := findNumInSet(connection[1], setList)

		if p1 == -1 && p2 == -1 {
			setList = append(setList, []int{connection[0], connection[1]})
			continue
		}

		if p1 == -1 {
			setList[p2] = append(setList[p2], connection[0])
			continue
		}

		if p2 == -1 {
			setList[p1] = append(setList[p1], connection[1])
			continue
		}

		//若p1 ！= p2,通过这条边就可以连接起来,将p2的数据挪到p1，同时删除p2
		if p1 != p2 {
			setList[p1] = append(setList[p1], setList[p2]...)
			setList = append(setList[:p2], setList[p2+1:]...)
			continue
		}

		//说明两端其实都在集合里，本来就是互通的，这条边是多余的
		invalidConnCnt++
	}

	//遍历找到所有的独立节点
	needConnNodeCnt := 0
	for num := 0; num < n; num++ {
		if findNumInSet(num, setList) == -1 {
			needConnNodeCnt++
		}
	}

	if len(setList) == 1 && needConnNodeCnt == 0 {
		return 0
	}

	//判断多余的边是否够用
	if invalidConnCnt >= (len(setList) + needConnNodeCnt - 1) {
		return len(setList) + needConnNodeCnt - 1
	}

	return -1
}

func findNumInSetV2(num int, setMap map[int]int) int {
	setNO, ok := setMap[num]
	if !ok {
		return -1
	}

	return setNO
}

//func MakeConnectedV2(n int, connections [][]int) int {
//	setMap := make(map[int]int)
//	setConnMap := make(map[int]int)
//
//	//遍历找到所有的互联小集合,同时找出多余线
//	invalidConnCnt := 0
//	setNO := 1
//	for _, connection := range connections {
//		p1 := findNumInSetV2(connection[0], setMap)
//		p2 := findNumInSetV2(connection[1], setMap)
//
//		if p1 == -1 && p2 == -1 {
//			setMap[connection[0]] = setNO
//			setMap[connection[1]] = setNO
//			setNO++
//			continue
//		}
//
//		if p1 == -1 {
//			setMap[connection[0]] = p2
//			continue
//		}
//
//		if p2 == -1 {
//			setMap[connection[1]] = p1
//			continue
//		}
//
//		//若p1 ！= p2,通过setConnMap存储两个集合互通这个关系
//		if p1 != p2 {
//			if p1 < p2 {
//				if _,ok := setConnMap[p2]
//				setConnMap[p2] = p1
//			} else {
//				setConnMap[p1] = p2
//			}
//			continue
//		}
//
//		//说明两端其实都在集合里，本来就是互通的，这条边是多余的
//		invalidConnCnt++
//	}
//
//	//遍历setMap
//	needConnNodeCnt := 0
//	realSetMap := make(map[int]bool)
//	for i := 0; i < n; i++ {
//		setInd := findNumInSetV2(i, setMap)
//		if setInd == -1 {
//			needConnNodeCnt++
//			continue
//		}
//
//		for {
//			data, ok := setConnMap[setInd]
//			if !ok {
//				break
//			}
//			setInd = data
//		}
//
//		realSetMap[setInd] = true
//	}
//
//	if len(realSetMap) == 1 && needConnNodeCnt == 0 {
//		return 0
//	}
//
//	//判断多余的边是否够用
//	if invalidConnCnt >= (len(realSetMap) + needConnNodeCnt - 1) {
//		return len(realSetMap) + needConnNodeCnt - 1
//	}
//
//	return -1
//}
