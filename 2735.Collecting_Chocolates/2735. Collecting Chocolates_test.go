package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{20, 1, 15},
				x:    5,
			},
			want: 13,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 3},
				x:    4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
