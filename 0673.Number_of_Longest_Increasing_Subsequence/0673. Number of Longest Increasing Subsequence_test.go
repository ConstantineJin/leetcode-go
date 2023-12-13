package main

import "testing"

func Test_findNumberOfLIS(t *testing.T) {
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
			args:    args{nums: []int{1, 3, 5, 4, 7}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 2, 2, 2, 2}},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findNumberOfLIS(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findNumberOfLIS() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
