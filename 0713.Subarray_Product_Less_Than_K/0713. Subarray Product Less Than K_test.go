package main

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
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
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			wantAns: 8,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 3},
				k:    0,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
