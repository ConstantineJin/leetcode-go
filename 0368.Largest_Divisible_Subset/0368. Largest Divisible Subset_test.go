package main

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3}},
			wantAns: []int{1, 2},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 4, 8}},
			wantAns: []int{1, 2, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
