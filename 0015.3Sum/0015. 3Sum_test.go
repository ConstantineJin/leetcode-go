package main

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
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
			args:    args{nums: []int{-1, 0, 1, 2, -1, -4}},
			wantAns: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 1, 1}},
			wantAns: nil,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{0, 0, 0}},
			wantAns: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := threeSum(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("threeSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
