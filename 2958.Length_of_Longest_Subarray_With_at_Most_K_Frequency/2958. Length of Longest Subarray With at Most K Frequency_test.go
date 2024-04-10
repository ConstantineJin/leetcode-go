package main

import "testing"

func Test_maxSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 1, 2},
				k:    2,
			},
			wantAns: 6,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2, 1, 2},
				k:    1,
			},
			wantAns: 2,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{5, 5, 5, 5, 5, 5, 5},
				k:    4,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxSubarrayLength(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxSubarrayLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
