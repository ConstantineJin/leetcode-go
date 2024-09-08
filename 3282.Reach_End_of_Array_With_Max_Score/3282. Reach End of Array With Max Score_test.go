package main

import "testing"

func Test_findMaximumScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 3, 1, 5}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{4, 3, 1, 3, 2}},
			wantAns: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMaximumScore(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findMaximumScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
