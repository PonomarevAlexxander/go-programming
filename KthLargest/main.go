package main

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return i > j
	})
	return nums[k-1]
}
