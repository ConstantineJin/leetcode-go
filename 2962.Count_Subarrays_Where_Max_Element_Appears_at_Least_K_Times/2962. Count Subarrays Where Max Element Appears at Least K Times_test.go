package main

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 3, 2, 3, 3},
				k:    2,
			},
			wantAns: 6,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 4, 2, 1},
				k:    3,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSubarrays(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
