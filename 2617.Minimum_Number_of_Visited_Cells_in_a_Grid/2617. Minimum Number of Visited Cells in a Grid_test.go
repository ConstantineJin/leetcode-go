package main

import "testing"

func Test_minimumVisitedCells(t *testing.T) {
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
			args: args{grid: [][]int{{3, 4, 2, 1}, {4, 2, 3, 1}, {2, 1, 0, 0}, {2, 4, 0, 0}}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{3, 4, 2, 1}, {4, 2, 1, 1}, {2, 1, 1, 0}, {3, 4, 1, 0}}},
			want: 3,
		},
		{
			name: "Example3",
			args: args{grid: [][]int{{2, 1, 0}, {1, 0, 0}}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumVisitedCells(tt.args.grid); got != tt.want {
				t.Errorf("minimumVisitedCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
