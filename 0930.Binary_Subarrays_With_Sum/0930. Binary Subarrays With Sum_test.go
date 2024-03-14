package main

import "testing"

func Test_numSubarraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 0, 1, 0, 1},
				goal: 2,
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{0, 0, 0, 0, 0},
				goal: 0,
			},
			wantAns: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSubarraysWithSum(tt.args.nums, tt.args.goal); gotAns != tt.wantAns {
				t.Errorf("numSubarraysWithSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
