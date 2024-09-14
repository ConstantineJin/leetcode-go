package main

import "testing"

func Test_longestSubarray(t *testing.T) {
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
			args:    args{nums: []int{1, 2, 3, 3, 2, 2}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3, 4}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubarray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("longestSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
