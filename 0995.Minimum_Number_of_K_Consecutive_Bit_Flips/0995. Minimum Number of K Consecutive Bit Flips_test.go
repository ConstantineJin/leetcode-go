package main

import "testing"

func Test_minKBitFlips(t *testing.T) {
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
				nums: []int{0, 1, 0},
				k:    1,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 0},
				k:    2,
			},
			wantAns: -1,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{0, 0, 0, 1, 0, 1, 1, 0},
				k:    3,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minKBitFlips(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("minKBitFlips() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
