package main

import "testing"

func Test_sumIndicesWithKSetBits(t *testing.T) {
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
				nums: []int{5, 10, 1, 5, 2},
				k:    1,
			},
			wantAns: 13,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{4, 3, 2, 1},
				k:    2,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumIndicesWithKSetBits(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("sumIndicesWithKSetBits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
