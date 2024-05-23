package main

import "testing"

func Test_longestEqualSubarray(t *testing.T) {
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
				nums: []int{1, 3, 2, 3, 1, 3},
				k:    3,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 2, 2, 1, 1},
				k:    2,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestEqualSubarray(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("longestEqualSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
