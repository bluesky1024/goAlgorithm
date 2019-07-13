package division

import "fmt"

/*问题*/
/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
说明:

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
不知道这道题为什么会在二分法里出现
但最简单的方式是用map判重
其次，将两个数组排序，定义两个ind，把2往1中数据上靠，这个过程也是o(m+n)复杂度
但排序是o(log(n)),优势是没有额外空间
*/
/*
v1是对应
nums1 [1,2,2,2,1] nums2[2,2,3]
res [2]
*/
func IntersectionV1(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)

	checkMap := make(map[int]bool)
	for _, v := range nums1 {
		checkMap[v] = true
	}

	for _, v := range nums2 {
		if checkRes, ok := checkMap[v]; ok && checkRes {
			checkMap[v] = false
			res = append(res, v)
		}
	}

	return res
}

/*
v2是对应
nums1 [1,2,2,2,1] nums2[2,2,3]
res [2,2]
*/
func IntersectionV2(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)

	checkMap := make(map[int]int)
	for _, v := range nums1 {
		if times, ok := checkMap[v]; ok {
			checkMap[v] = times + 1
		} else {
			checkMap[v] = 1
		}
	}

	checkTimesMap := make(map[int]int)
	for _, v := range nums2 {
		if times, ok := checkMap[v]; ok {
			if repeatTimes, ok := checkTimesMap[v]; ok {
				if repeatTimes < times {
					checkTimesMap[v] = repeatTimes + 1
					res = append(res, v)
				}
			} else {
				checkTimesMap[v] = 1
				res = append(res, v)
			}
		}
	}

	return res
}

/*问题*/
/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1


提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
典型的二分法。。。
*/
func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	head := 0
	tail := length - 1
	if nums[head] == target {
		return head
	}
	if nums[tail] == target {
		return tail
	}
	for {
		fmt.Println("a", head, tail)
		if nums[head] > target {
			return -1
		}
		if nums[tail] < target {
			return -1
		}
		if head+1 >= tail {
			break
		}

		mid := (head + tail) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			head = mid
		} else {
			tail = mid
		}
	}
	return -1
}
