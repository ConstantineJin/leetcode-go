package main

import "testing"

func Test_minimumMoves(t *testing.T) {
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
			args: args{grid: [][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}}},
			want: 3,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{1, 3, 0}, {1, 0, 0}, {1, 0, 3}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoves(tt.args.grid); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
