package sort

import "fmt"

func SelectSort(nums []int) []int {
	length := len(nums)
	cnt1 := 0
	cnt2 := 0
	for i := 0; i < length; i++ {
		maxInd := i
		tempMax := nums[i]
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

func QuickSort(nums []int, sInd int, eInd int) {

}

func MergeSort(nums []int) {

}
