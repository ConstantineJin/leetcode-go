package main

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}},
			wantAns: []int{3, 2},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 1, 2}},
			wantAns: []int{1},
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1}},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findDuplicates(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findDuplicates() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
