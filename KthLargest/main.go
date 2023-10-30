// Package provides function for finding Kth largest/smallest element using quick select
// Complexity avg case O(n)
// Naming and no errors handling are caused by CI tests provided by course
package main

func findKthLargest(nums []int, k int) int {
	return findKthSmallest(nums, len(nums)-k+1)
}

func findKthSmallest(nums []int, k int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	return doFindKthSmallest(numsCopy, k)
}

func doFindKthSmallest(nums []int, k int) int {
	pivotIndex := partition(nums)
	if pivotIndex+1 == k {
		return nums[pivotIndex]
	}
	if pivotIndex+1 < k {
		return doFindKthSmallest(nums[pivotIndex+1:], k-pivotIndex-1)
	}
	return doFindKthSmallest(nums[:pivotIndex], k)
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
