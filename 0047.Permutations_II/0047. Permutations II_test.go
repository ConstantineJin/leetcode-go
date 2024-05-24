package main

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
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
			args:    args{nums: []int{1, 1, 2}},
			wantAns: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := permuteUnique(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("permuteUnique() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
