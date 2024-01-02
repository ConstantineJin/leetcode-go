package main

import (
	"reflect"
	"testing"
)

func Test_findMatrix(t *testing.T) {
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
			args: args{nums: []int{1, 3, 4, 1, 2, 3, 1}},
			want: [][]int{{1, 2, 3, 4}, {1, 3}, {1}},
		},
		{
			name: "Example2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: [][]int{{1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatrix(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
