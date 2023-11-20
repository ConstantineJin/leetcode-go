package main

import "testing"

func Test_findShortestSubArray(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 2, 3, 1}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 2, 3, 1, 4, 2}},
			wantAns: 6,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{1, 2, 1}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findShortestSubArray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findShortestSubArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
