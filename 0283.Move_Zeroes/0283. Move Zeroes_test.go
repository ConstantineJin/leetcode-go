package main

import (
	"reflect"
	"testing"
)

func testMoveZeroes(nums []int) []int {
	moveZeroes(nums)
	return nums
}

func Test_testMoveZeroes(t *testing.T) {
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
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Example2",
			args: args{nums: []int{0}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testMoveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testMoveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
