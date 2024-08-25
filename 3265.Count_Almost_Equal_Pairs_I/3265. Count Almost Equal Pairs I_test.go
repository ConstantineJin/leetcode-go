package main

import "testing"

func Test_countPairs(t *testing.T) {
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
			args:    args{nums: []int{3, 12, 30, 17, 21}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 1, 1, 1, 1}},
			wantAns: 10,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{123, 231}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPairs(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
