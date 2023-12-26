package main

import (
	"reflect"
	"testing"
)

func testNextPermutation(nums []int) []int {
	nextPermutation(nums)
	return nums
}

func Test_testNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
		{
			name: "Example2",
			args: args{nums: []int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "Example3",
			args: args{nums: []int{1, 1, 5}},
			want: []int{1, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testNextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testNextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
