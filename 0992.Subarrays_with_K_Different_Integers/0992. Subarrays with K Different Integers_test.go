package main

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
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
				nums: []int{1, 2, 1, 2, 3},
				k:    2,
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 1, 3, 4},
				k:    3,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subarraysWithKDistinct(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
