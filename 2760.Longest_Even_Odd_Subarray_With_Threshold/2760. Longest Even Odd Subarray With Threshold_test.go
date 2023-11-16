package main

import "testing"

func Test_longestAlternatingSubarray(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{3, 2, 5, 4}, threshold: 5},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2}, threshold: 2},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{nums: []int{2, 3, 4, 5}, threshold: 4},
			wantAns: 3,
		},
		{
			name:    "Example4",
			args:    args{nums: []int{4, 10, 3}, threshold: 10},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestAlternatingSubarray(tt.args.nums, tt.args.threshold); gotAns != tt.wantAns {
				t.Errorf("longestAlternatingSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
