package main

import "testing"

func Test_maxOperations(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
				k:    5,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{3, 1, 3, 4, 3},
				k:    6,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxOperations(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
