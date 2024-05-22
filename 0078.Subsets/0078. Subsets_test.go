package main

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{nil, {3}, {2}, {2, 3}, {1}, {1, 3}, {1, 2}, {1, 2, 3}},
		},
		{
			name: "Example2",
			args: args{nums: []int{0}},
			want: [][]int{nil, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
