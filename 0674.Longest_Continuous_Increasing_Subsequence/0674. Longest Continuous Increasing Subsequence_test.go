package main

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
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
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{2, 2, 2, 2, 2}},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1, 3, 5, 7}},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findLengthOfLCIS(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findLengthOfLCIS() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
