package main

import (
	"reflect"
	"testing"
)

func testSortColors(nums []int) []int {
	sortColors(nums)
	return nums
}

func Test_testSortColors(t *testing.T) {
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
			args: args{nums: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Example2",
			args: args{nums: []int{2, 0, 1}},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testSortColors(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testSortColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
