package main

import "testing"

func Test_minOperations(t *testing.T) {
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
				nums: []int{2, 11, 10, 1, 3},
				k:    10,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 2, 4, 9},
				k:    20,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minOperations(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("minOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
