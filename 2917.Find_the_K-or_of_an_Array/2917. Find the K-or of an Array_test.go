package main

import "testing"

func Test_findKOr(t *testing.T) {
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
				nums: []int{7, 12, 9, 8, 9, 15},
				k:    4,
			},
			wantAns: 9,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{2, 12, 1, 11, 4, 5},
				k:    6,
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{10, 8, 5, 9, 11, 6, 8},
				k:    1,
			},
			wantAns: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findKOr(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("findKOr() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
