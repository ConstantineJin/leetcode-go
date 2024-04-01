package main

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		minK int
		maxK int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 3, 5, 2, 7, 5},
				minK: 1,
				maxK: 5,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 1, 1},
				minK: 1,
				maxK: 1,
			},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSubarrays(tt.args.nums, tt.args.minK, tt.args.maxK); gotAns != tt.wantAns {
				t.Errorf("countSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
