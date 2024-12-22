package main

import (
	"reflect"
	"testing"
)

func Test_minSubsequence(t *testing.T) {
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
			args:    args{nums: []int{4, 3, 10, 9, 8}},
			wantAns: []int{10, 9},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{4, 4, 7, 6, 7}},
			wantAns: []int{7, 7, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minSubsequence(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("minSubsequence() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
