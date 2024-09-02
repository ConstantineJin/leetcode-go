package main

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{grid: [][]int{{1, 2, 3}, {4, 3, 2}, {1, 1, 1}}},
			want: 8,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{8, 7, 6}, {8, 3, 2}}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.grid); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
