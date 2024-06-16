package main

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 3},
				n:    6,
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 5, 10},
				n:    20,
			},
			wantAns: 2,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{1, 2, 2},
				n:    5,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minPatches(tt.args.nums, tt.args.n); gotAns != tt.wantAns {
				t.Errorf("minPatches() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
