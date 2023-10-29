package main

func findKthLargest(nums []int, k int) int {
	return findKthSmallest(nums, len(nums)-k+1)
}

func findKthSmallest(nums []int, k int) int {
	pivotIndex := partition(nums)
	if pivotIndex+1 == k {
		return nums[pivotIndex]
	}
	if pivotIndex+1 < k {
		return findKthSmallest(nums[pivotIndex+1:], k-pivotIndex-1)
	}
	return findKthSmallest(nums[:pivotIndex], k)
}

func partition(nums []int) int {
	pivot := nums[len(nums)-1]
	pivotIndex := 0
	for index, element := range nums {
		if element < pivot {
			nums[pivotIndex], nums[index] = nums[index], nums[pivotIndex]
			pivotIndex++
		}
	}
	nums[pivotIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[pivotIndex]
	return pivotIndex
}
