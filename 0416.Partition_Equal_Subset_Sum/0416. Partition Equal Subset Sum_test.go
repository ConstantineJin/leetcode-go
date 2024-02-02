package main

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns bool
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 5, 11, 5}},
			wantAns: true,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3, 5}},
			wantAns: false,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{14, 9, 8, 4, 3, 2}},
			wantAns: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := canPartition(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("canPartition() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
