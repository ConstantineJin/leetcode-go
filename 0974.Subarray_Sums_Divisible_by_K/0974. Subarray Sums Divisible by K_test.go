package main

import "testing"

func Test_subarraysDivByK(t *testing.T) {
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
				nums: []int{4, 5, 0, -2, -3, 1},
				k:    5,
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{5},
				k:    9,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subarraysDivByK(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("subarraysDivByK() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
