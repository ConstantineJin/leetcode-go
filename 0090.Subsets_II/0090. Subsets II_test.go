package main

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 2}},
			wantAns: [][]int{{1, 2, 2}, {1, 2}, {1}, {2, 2}, {2}, {}},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0}},
			wantAns: [][]int{{0}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("subsetsWithDup() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
