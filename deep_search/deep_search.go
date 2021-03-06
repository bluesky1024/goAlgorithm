package deep_search

import (
	"fmt"
	"github.com/bluesky1024/goAlgorithm/tree"
	"math"
	"strconv"
	"sync"
)

/*问题*/
/*
有 N 个房间，开始时你位于 0 号房间。每个房间有不同的号码：0，1，2，...，N-1，并且房间里可能有一些钥匙能使你进入下一个房间。

在形式上，对于每个房间 i 都有一个钥匙列表 rooms[i]，每个钥匙 rooms[i][j] 由 [0,1，...，N-1] 中的一个整数表示，其中 N = rooms.length。 钥匙 rooms[i][j] = v 可以打开编号为 v 的房间。

最初，除 0 号房间外的其余所有房间都被锁住。

你可以自由地在房间之间来回走动。

如果能进入每个房间返回 true，否则返回 false。

示例 1：

输入: [[1],[2],[3],[]]
输出: true
解释:
我们从 0 号房间开始，拿到钥匙 1。
之后我们去 1 号房间，拿到钥匙 2。
然后我们去 2 号房间，拿到钥匙 3。
最后我们去了 3 号房间。
由于我们能够进入每个房间，我们返回 true。
示例 2：

输入：[[1,3],[3,0,1],[2],[0]]
输出：false
解释：我们不能进入 2 号房间。
提示：

1 <= rooms.length <= 1000
0 <= rooms[i].length <= 1000
所有房间中的钥匙数量总计不超过 3000。
*/
/*思路*/
/*
沿着一个房间的所有选择进行遍历
走过的房间不要再走
*/
func searchDeepRooms(curRoom int, rooms [][]int, lock []bool) {
	lock[curRoom] = true
	for _, v := range rooms[curRoom] {
		if !lock[v] {
			searchDeepRooms(v, rooms, lock)
		}
	}
}
func CanVisitAllRooms(rooms [][]int) bool {
	lock := make([]bool, len(rooms))
	searchDeepRooms(0, rooms, lock)
	for _, v := range lock {
		if !v {
			return false
		}
	}
	return true
}

/*问题*/
/*
视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。

给定一个字符串 ring，表示刻在外环上的编码；给定另一个字符串 key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。

最初，ring 的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。

旋转 ring 拼出 key 字符 key[i] 的阶段中：

您可以将 ring 顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 key[i] 。
如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段）, 直至完成所有拼写。
示例：
图片省略：一个圆圈，中间一个button，四周圆环上按顺序分布字符ring
输入: ring = "godding", key = "gd"
输出: 4
解释:
 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 当然, 我们还需要1步进行拼写。
 因此最终的输出是 4。
*/
/*思路*/
/*
同样深度优先进行遍历，遍历到最后一个字符再返回结果
取步数最少的结果
*/
func chooseOneStep(ringLength int, keyLength int, ring string, key string, curPos int, curInd, curStepCnt int, curMin *int) {
	fmt.Println(curInd)
	if ring[curPos] == key[curInd] {
		if curInd == keyLength-1 {
			if *curMin == -1 {
				*curMin = curStepCnt
			} else {
				if *curMin > curStepCnt {
					*curMin = curStepCnt
				}
			}
		} else {
			chooseOneStep(ringLength, keyLength, ring, key, curPos, curInd+1, curStepCnt, curMin)
		}
		return
	}
	wg := sync.WaitGroup{}
	for i := 0; i < ringLength; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if ring[i] == key[curInd] {
				tempStep1 := math.Abs(float64(curPos - i))
				tempStep2 := float64(ringLength) - math.Abs(float64(curPos-i))
				tempCurStepCnt := curStepCnt + int(math.Min(tempStep1, tempStep2))
				if curInd == keyLength-1 {
					if *curMin == -1 {
						*curMin = tempCurStepCnt
					} else {
						if *curMin > tempCurStepCnt {
							fmt.Println(tempCurStepCnt)
							*curMin = tempCurStepCnt
						}
					}
				} else {
					chooseOneStep(ringLength, keyLength, ring, key, i, curInd+1, tempCurStepCnt, curMin)
				}
			}
		}(i)
	}
	wg.Wait()
}
func FindRotateSteps(ring string, key string) int {
	ringLength := len(ring)
	keyLength := len(key)
	curMin := -1
	chooseOneStep(ringLength, keyLength, ring, key, 0, 0, 0, &curMin)
	curMin += keyLength
	return curMin
}

/*问题*/
/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
父节点的路径集合是两个子节点的路径集合与父节点本身拼凑后相加
*/
func BinaryTreePaths(root *tree.TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	lSlice := BinaryTreePaths(root.Left)
	rSlice := BinaryTreePaths(root.Right)
	res := make([]string, 0)
	for _, s := range lSlice {
		res = append(res, strconv.Itoa(root.Val)+"->"+s)
	}
	for _, s := range rSlice {
		res = append(res, strconv.Itoa(root.Val)+"->"+s)
	}
	return res
}

/*问题*/
/*
有 N 个网络节点，标记为 1 到 N。

给定一个列表 times，表示信号经过有向边的传递时间。 times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w 是一个信号从源节点传递到目标节点的时间。

现在，我们从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。



示例：
		2

	1


输入：times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
输出：2


注意:

N 的范围在 [1, 100] 之间。
K 的范围在 [1, N] 之间。
times 的长度在 [1, 6000] 之间。
所有的边 times[i] = (u, v, w) 都有 1 <= u, v <= N 且 0 <= w <= 100。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/network-delay-time
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
1.需要将原有的times [][]int结构改成map[startNode][][]int，方便后续逐个node定位数据
2.从指定节点出发，记录可达节点，及其当前所需时间（多种可能中的最短时间）
3.深度遍历到最后一级，没有新的节点可以延伸了，就停止
*/
func NetworkDelayTime(times [][]int, N int, K int) int {
	relations := make(map[int][][]int)
	for _, time := range times {
		relations[time[0]] = append(relations[time[0]], []int{time[1], time[2]})
	}

	nodesPass := make(map[int]int)
	nodesPass[K] = 0
	searchNetworkNode(relations, nodesPass, K, 0)

	maxTime := 0
	for i := 1; i <= N; i++ {
		time, ok := nodesPass[i]
		if !ok {
			return -1
		}
		if maxTime < time {
			maxTime = time
		}
	}

	return maxTime
}

func searchNetworkNode(relations map[int][][]int, nodesPass map[int]int, curNode int, curTime int) {
	curNodeRelation, ok := relations[curNode]
	if !ok {
		return
	}

	for _, relation := range curNodeRelation {
		nextNode := relation[0]

		passData, ok := nodesPass[nextNode]
		if !ok || passData > curTime+relation[1] {
			nodesPass[nextNode] = curTime + relation[1]
			searchNetworkNode(relations, nodesPass, nextNode, curTime+relation[1])
			continue
		}
	}
}
