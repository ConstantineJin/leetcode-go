package main

import "testing"

func Test_maxScore(t *testing.T) {
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
				nums: []int{2, 3, 6, 1, 9, 2},
				x:    5,
			},
			want: 13,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{2, 4, 6, 8},
				x:    3,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
