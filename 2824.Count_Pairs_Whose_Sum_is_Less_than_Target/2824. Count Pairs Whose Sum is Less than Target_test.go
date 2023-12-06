package main

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums:   []int{-1, 1, 2, 3, 1},
				target: 2,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				nums:   []int{-6, 2, 5, -2, -7, -1, 3},
				target: -2,
			},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPairs(tt.args.nums, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("countPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
