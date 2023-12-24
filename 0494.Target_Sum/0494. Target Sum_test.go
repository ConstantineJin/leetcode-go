package main

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
		{
			name: "Example2",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
