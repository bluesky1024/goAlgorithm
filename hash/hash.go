package hash

import (
	"github.com/bluesky1024/goAlgorithm/sort"
	"strconv"
)

/*问题*/
/*
在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。

返回重复了 N 次的那个元素。



示例 1：

输入：[1,2,3,3]
输出：3
示例 2：

输入：[2,1,2,5,3,2]
输出：2
示例 3：

输入：[5,1,5,2,5,3,5,4]
输出：5


提示：

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length 为偶数
*/
/*思路*/
/*
N+1个元素，有一个元素重复N次
说明其他元素只有一次
说明如果有个元素重复了，那一定就是目标元素
*/
func RepeatedNTimes(A []int) int {
	tempMap := make(map[int]bool, len(A)/2)
	for _, v := range A {
		_, ok := tempMap[v]
		if ok {
			return v
		} else {
			tempMap[v] = true
		}
	}
	return 0
}

/*问题*/
/*
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。

h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”



示例:

输入: citations = [3,0,6,1,5]
输出: 3
解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。


说明: 如果 h 有多种可能的值，h 指数是其中最大的那个。
*/
/*思路*/
/*
至多有h篇文章分别被至少引用了h次（被引用次数为100的文章，也可以看作至少被引用了一次）
*/
func HIndex(citations []int) int {
	tempRes := make(map[int]int, len(citations))
	for _, v1 := range citations {
		if v1 == 0 {
			continue
		}
		temp := 0
		for _, v2 := range citations {
			if v1 <= v2 {
				temp++
			}
		}
		tempRes[v1] = temp
	}

	maxH := 0
	for k, v := range tempRes {
		min := k
		if k > v {
			min = v
		}
		if maxH < min {
			maxH = min
		}
	}
	return maxH
}

/*问题*/
/*
给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。

网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。



示例 :

输入:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

输出: 16

解释: 它的周长是下面图片中的 16 个黄色的边：
*/
/*思路*/
/*
宽高都不超过100,map[int]bool存储陆地，key=长*1000+宽，val=1
一个"1"点周围有上下左右4条边，如果某方向临近"1"，则不算在周长内
遍历map来统计周长
判断周围是否有边，直接根据map是否存在来判断

简直思维僵化，要什么map,直接用grid就能定位。。。
*/
func IslandPerimeter(grid [][]int) int {
	lenRow := len(grid)
	if lenRow == 0 {
		return 0
	}
	lenCol := len(grid[0])
	if lenCol == 0 {
		return 0
	}

	posMap := make(map[int]bool)
	for i, l := range grid {
		for j, v := range l {
			if v == 1 {
				posMap[i*1000+j] = true
			}
		}
	}

	res := 0
	for key, _ := range posMap {
		row := key / 1000
		col := key % 1000

		if row == 0 {
			res++
		} else {
			if _, ok := posMap[(row-1)*1000+col]; !ok {
				res++
			}
		}

		if row == lenRow-1 {
			res++
		} else {
			if _, ok := posMap[(row+1)*1000+col]; !ok {
				res++
			}
		}

		if col == 0 {
			res++
		} else {
			if _, ok := posMap[row*1000+col-1]; !ok {
				res++
			}
		}

		if col == lenCol-1 {
			res++
		} else {
			if _, ok := posMap[row*1000+col+1]; !ok {
				res++
			}
		}
	}
	return res
}

/*问题*/
/*
我们有一个项的集合，其中第 i 项的值为 values[i]，标签为 labels[i]。

我们从这些项中选出一个子集 S，这样一来：

|S| <= num_wanted
对于任意的标签 L，子集 S 中标签为 L 的项的数目总满足 <= use_limit。
返回子集 S 的最大可能的 和。



示例 1：

输入：values = [5,4,3,2,1], labels = [1,1,2,2,3], num_wanted = 3, use_limit = 1
输出：9
解释：选出的子集是第一项，第三项和第五项。
示例 2：

输入：values = [5,4,3,2,1], labels = [1,3,3,3,2], num_wanted = 3, use_limit = 2
输出：12
解释：选出的子集是第一项，第二项和第三项。
示例 3：

输入：values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 1
输出：16
解释：选出的子集是第一项和第四项。
示例 4：

输入：values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 2
输出：24
解释：选出的子集是第一项，第二项和第四项。


提示：

1 <= values.length == labels.length <= 20000
0 <= values[i], labels[i] <= 20000
1 <= num_wanted, use_limit <= values.length
*/
/*思路*/
/*
1. 纯暴力遍历
2. 想想。。。
返回最大的可能的和
先把value排序
优先把大的数放进去
放进去之前先看看是否超过label数量限制了

之前以为values是集合不会重复，现在需要考虑重复数据
那么问题来了，怎么存储重复的数字的不同label==>valLabelMap
遍历过了的数字的次数存储方式=>valTimesMap
遍历过来了的标签的次数存储方式=>labelTimesMap
*/

func LargestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	if use_limit == 0 || num_wanted == 0 {
		return 0
	}
	valLabelMap := make(map[int][]int)
	for i, v := range values {
		valLabelMap[v] = append(valLabelMap[v], labels[i])
	}
	//从大到小排序
	sort.QuickSort(values)

	resArr := make([]int, 0)
	valTimesMap := make(map[int]int)
	labelTimesMap := make(map[int]int)
	for _, v := range values {
		if len(resArr) == num_wanted {
			break
		}

		valTime, vok := valTimesMap[v]
		if !vok {
			valTime = 0
			valTimesMap[v] = 1
		} else {
			valTimesMap[v] = valTime + 1
		}

		time, ok := labelTimesMap[valLabelMap[v][valTime]]
		if !ok {
			resArr = append(resArr, v)
			labelTimesMap[valLabelMap[v][valTime]] = 1
			continue
		}

		if time >= use_limit {
			continue
		}

		resArr = append(resArr, v)
		labelTimesMap[valLabelMap[v][valTime]] = time + 1
	}

	res := 0
	for _, v := range resArr {
		res = res + v
	}
	return res
}

/*问题*/
/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
/*思路*/
/*
似乎只能先排序，后暴力
想办法尽量减少循环
不能重复，所以进入res之前需要先于上一个做比较
*/
func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	checkMap := make(map[string]bool)

	length := len(nums)
	if length < 4 {
		return res
	}
	sort.QuickSort(nums)

	var curV int
	for i := 0; i < length-3; i++ {
		curV = 0
		if (nums[i] + nums[i+1] + nums[i+2] + nums[i+3]) < target {
			break
		}
		curV = curV + nums[i]
		for j := i + 1; j < length-2; j++ {
			if (curV + nums[j] + nums[j+1] + nums[j+2]) < target {
				break
			}
			curV = curV + nums[j]
			for k := j + 1; k < length-1; k++ {
				if (curV + nums[k] + nums[k+1]) < target {
					break
				}
				curV = curV + nums[k]
				for l := k + 1; l < length; l++ {
					if curV+nums[l] < target {
						break
					}
					curV = curV + nums[l]
					if curV == target {
						temp := []int{nums[i], nums[j], nums[k], nums[l]}
						mapKey := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[k]) + strconv.Itoa(nums[l])
						if _, ok := checkMap[mapKey]; !ok {
							res = append(res, temp)
							checkMap[mapKey] = true
						}
					}
					curV = curV - nums[l]
				}
				curV = curV - nums[k]
			}
			curV = curV - nums[j]
		}
	}

	return res
}
