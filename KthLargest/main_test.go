package main

import (
	"testing"
)

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple test", args{[]int{7, 2, 3, 8, 0, 5}, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.args.nums))
			copy(numsCopy, tt.args.nums)
			if got := findKthSmallest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
			for index := range tt.args.nums {
				if tt.args.nums[index] != numsCopy[index] {
					t.Errorf("findKthSmallest() = %v, want %v", tt.args.nums, numsCopy)
				}
			}
		})
	}
}
