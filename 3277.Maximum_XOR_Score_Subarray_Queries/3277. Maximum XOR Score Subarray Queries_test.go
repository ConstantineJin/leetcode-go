package main

import (
	"reflect"
	"testing"
)

func Test_maximumSubarrayXor(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:    []int{2, 8, 4, 32, 16, 1},
				queries: [][]int{{0, 2}, {1, 4}, {0, 5}},
			},
			want: []int{12, 60, 60},
		},
		{
			name: "Example2",
			args: args{
				nums:    []int{0, 7, 3, 2, 8, 5, 1},
				queries: [][]int{{0, 3}, {1, 5}, {2, 4}, {2, 6}, {5, 6}},
			},
			want: []int{7, 14, 11, 14, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarrayXor(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumSubarrayXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
