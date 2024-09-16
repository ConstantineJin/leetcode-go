package main

import (
	"reflect"
	"testing"
)

func Test_getSneakyNumbers(t *testing.T) {
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
			args:    args{nums: []int{0, 1, 1, 0}},
			wantAns: []int{1, 0},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 3, 2, 1, 3, 2}},
			wantAns: []int{3, 2},
		},
		{
			name:    "Example3",
			args:    args{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}},
			wantAns: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getSneakyNumbers(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getSneakyNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
