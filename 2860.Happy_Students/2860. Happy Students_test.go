package main

import (
	"testing"
)

func Test_countWays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 1}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{6, 0, 3, 3, 6, 7, 2, 7}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countWays(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countWays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
