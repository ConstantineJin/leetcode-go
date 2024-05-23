package main

import "testing"

func Test_beautifulSubsets(t *testing.T) {
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
				nums: []int{2, 4, 6},
				k:    2,
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := beautifulSubsets(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("beautifulSubsets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
