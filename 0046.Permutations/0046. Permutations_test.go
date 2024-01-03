package main

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 1}},
			wantAns: [][]int{{0, 1}, {1, 0}},
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1}},
			wantAns: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := permute(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("permute() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
