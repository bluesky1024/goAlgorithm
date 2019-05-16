package deep_search

import (
	"math"
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
	//if ring[curPos] == key[curInd] {
	//	if curInd == keyLength-1 {
	//		if *curMin == -1 {
	//			*curMin = curStepCnt
	//		} else {
	//			if *curMin > curStepCnt {
	//				*curMin = curStepCnt
	//			}
	//		}
	//	} else {
	//		curInd++
	//		chooseOneStep(ringLength, keyLength, ring, key, curPos, curInd, curStepCnt, curMin)
	//	}
	//	return
	//}
	for i := 0; i < ringLength; i++ {
		if ring[i] == key[curInd] {
			tempStep1 := math.Abs(float64(curPos - i))
			tempStep2 := float64(ringLength) - math.Abs(float64(curPos-i))
			curPos = i
			//if tempStep1 <= tempStep2 {
			//	curStepCnt += tempStep1
			//} else {
			//	curStepCnt += tempStep2
			//}
			curStepCnt += int(math.Min(tempStep1, tempStep2))
			if curInd == keyLength-1 {
				if *curMin == -1 {
					*curMin = curStepCnt
				} else {
					if *curMin > curStepCnt {
						*curMin = curStepCnt
					}
				}
			} else {
				curInd++
				chooseOneStep(ringLength, keyLength, ring, key, curPos, curInd, curStepCnt, curMin)
			}
		}
	}
}
func FindRotateSteps(ring string, key string) int {
	ringLength := len(ring)
	keyLength := len(key)
	curMin := -1
	chooseOneStep(ringLength, keyLength, ring, key, 0, 0, 0, &curMin)
	curMin += keyLength
	return curMin
}
