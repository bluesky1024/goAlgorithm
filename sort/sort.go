package sort

import (
	"fmt"
	origSort "sort"
)

func SelectSort(nums []int) []int {
	length := len(nums)
	cnt1 := 0
	cnt2 := 0
	for i := 0; i < length; i++ {
		maxInd := i
		tempMax := nums[i]
		fmt.Println(nums)
		for j := i + 1; j < length; j++ {
			cnt2++
			if tempMax < nums[j] {
				maxInd = j
				tempMax = nums[j]
			}
		}
		cnt1++
		nums[i], nums[maxInd] = nums[maxInd], nums[i]
	}
	fmt.Println(cnt1, cnt2)
	return nums
}

func BubbleSort(nums []int) []int {
	length := len(nums)
	ind := length
	cnt1 := 0
	cnt2 := 0
	for i := 0; i < length; i++ {
		fmt.Println(nums)
		for j := 0; j < ind-1; j++ {
			cnt2++
			if nums[j] < nums[j+1] {
				cnt1++
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		ind--
	}
	fmt.Println(cnt1, cnt2)
	return nums
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid, i := nums[0], 1
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[i] < mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	nums[head] = mid
	QuickSort(nums[:head])
	QuickSort(nums[head+1:])
}

/*
归并排序是稳定排序
时间复杂度O(nlogn)。为什么这么快？
n/2 + n/2 merge to n , 比较n次
分组粒度到1，n可分为logn/log2组
4->2 + 2->1+1 + 1+1
*/
func MergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	fmt.Println(num)
	left := MergeSort(r[:num])
	right := MergeSort(r[num:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	fmt.Println(left, right)
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

/*堆排序*/
func HeapSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	for i := 0; i < length; i++ {
		minAjust(nums[i:])
	}
}

func minAjust(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := length/2 - 1; i >= 0; i-- {
		if (2*i+1 <= length-1) && (nums[i] >= nums[2*i+1]) {
			nums[i], nums[2*i+1] = nums[2*i+1], nums[i]
		}
		if (2*i+2 <= length-1) && (nums[i] >= nums[2*i+2]) {
			nums[i], nums[2*i+2] = nums[2*i+2], nums[i]
		}
	}
}

/*问题*/
/*
给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。

如果不能形成任何面积不为零的三角形，返回 0。



示例 1：

输入：[2,1,2]
输出：5
示例 2：

输入：[1,2,1]
输出：0
示例 3：

输入：[3,2,3,4]
输出：10
示例 4：

输入：[3,6,2,3]
输出：8


提示：

3 <= A.length <= 10000
1 <= A[i] <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-perimeter-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
构成三角形的条件，必须任意两条边的和大于第三条边
将数组从大到小排序，从最大值开始，只要第一个值>后两个值的和就满足条件了，优先最大值
*/
func LargestPerimeter(A []int) int {
	length := len(A)
	if length < 3 {
		return 0
	}
	res := 0
	origSort.Ints(A)
	for i := length - 1; i >= 2; i-- {
		if A[i] < (A[i-1] + A[i-2]) {
			res = A[i] + A[i-1] + A[i-2]
			break
		}
	}
	return res
}

/*问题*/
/*
给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行）以完成对数组 A 的排序。

返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。



示例 1：

输入：[3,2,4,1]
输出：[4,2,4,3]
解释：
我们执行 4 次煎饼翻转，k 值分别为 4，2，4，和 3。
初始状态 A = [3, 2, 4, 1]
第一次翻转后 (k=4): A = [1, 4, 2, 3]
第二次翻转后 (k=2): A = [4, 1, 2, 3]
第三次翻转后 (k=4): A = [3, 2, 1, 4]
第四次翻转后 (k=3): A = [1, 2, 3, 4]，此时已完成排序。
示例 2：

输入：[1,2,3]
输出：[]
解释：
输入已经排序，因此不需要翻转任何内容。
请注意，其他可能的答案，如[3，3]，也将被接受。


提示：

1 <= A.length <= 100
A[i] 是 [1, 2, ..., A.length] 的排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pancake-sorting
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
没找到最高效的方法，但以下方法能固定得到正确顺序
1.每i次遍历，定位前length-i位最大值位置n
2.先根据n进行反转将最大值挪到首位，再根据i反转，将最大值挪到第length-i位
3.遍历直至只剩最后一个数
*/
func PancakeSort(A []int) []int {
	length := len(A)
	if length <= 1 {
		return nil
	}
	res := make([]int, 2)
	maxInd := -1
	max := -1
	for i, v := range A {
		if v > max {
			maxInd = i
			max = v
		}
	}
	if maxInd == length-1 {
		return PancakeSort(A[:length-1])
	}
	res[0] = maxInd + 1
	for i := 0; i <= maxInd/2; i++ {
		A[i], A[maxInd-i] = A[maxInd-i], A[i]
	}
	res[1] = length

	//逆转位置
	for i := 0; i < length/2; i++ {
		A[i], A[length-i-1] = A[length-i-1], A[i]
	}

	left := PancakeSort(A[:length-1])
	return append(res, left...)
}

/*问题*/
/*
给你一个整数数组 arr 和一个整数 k 。

设 m 为数组的中位数，只要满足下述两个前提之一，就可以判定 arr[i] 的值比 arr[j] 的值更强：

 |arr[i] - m| > |arr[j] - m|
 |arr[i] - m| == |arr[j] - m|，且 arr[i] > arr[j]
请返回由数组中最强的 k 个值组成的列表。答案可以以 任意顺序 返回。

中位数 是一个有序整数列表中处于中间位置的值。形式上，如果列表的长度为 n ，那么中位数就是该有序列表（下标从 0 开始）中位于 ((n - 1) / 2) 的元素。

例如 arr = [6, -3, 7, 2, 11]，n = 5：数组排序后得到 arr = [-3, 2, 6, 7, 11] ，数组的中间位置为 m = ((5 - 1) / 2) = 2 ，中位数 arr[m] 的值为 6 。
例如 arr = [-7, 22, 17, 3]，n = 4：数组排序后得到 arr = [-7, 3, 17, 22] ，数组的中间位置为 m = ((4 - 1) / 2) = 1 ，中位数 arr[m] 的值为 3 。


示例 1：

输入：arr = [1,2,3,4,5], k = 2
输出：[5,1]
解释：中位数为 3，按从强到弱顺序排序后，数组变为 [5,1,4,2,3]。最强的两个元素是 [5, 1]。[1, 5] 也是正确答案。
注意，尽管 |5 - 3| == |1 - 3| ，但是 5 比 1 更强，因为 5 > 1 。
示例 2：

输入：arr = [1,1,3,5,5], k = 2
输出：[5,5]
解释：中位数为 3, 按从强到弱顺序排序后，数组变为 [5,5,1,1,3]。最强的两个元素是 [5, 5]。
示例 3：

输入：arr = [6,7,11,7,6,8], k = 5
输出：[11,8,6,6,7]
解释：中位数为 7, 按从强到弱顺序排序后，数组变为 [11,8,6,6,7,7]。
[11,8,6,6,7] 的任何排列都是正确答案。
示例 4：

输入：arr = [6,-3,7,2,11], k = 3
输出：[-3,11,2]
示例 5：

输入：arr = [-7,22,17,3], k = 2
输出：[22,17]


提示：

1 <= arr.length <= 10^5
-10^5 <= arr[i] <= 10^5
1 <= k <= arr.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/the-k-strongest-values-in-an-array
*/
/*思路*/
/*
先排序，找到中位数
再从两端开始按照规则遍历，找出前k个数
*/
func GetStrongest(arr []int, k int) []int {
	if len(arr) == 1 {
		return arr
	}

	origSort.Ints(arr)

	mid := (len(arr) - 1) / 2
	left := 0
	right := len(arr) - 1

	res := make([]int, 0)
	for {
		if len(res) == k {
			break
		}
		if arr[mid]-arr[left] > arr[right]-arr[mid] {
			res = append(res, arr[left])
			left++
		} else {
			res = append(res, arr[right])
			right--
		}
	}

	return res
}
