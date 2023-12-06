package main

import "testing"

func Test_minPathCost(t *testing.T) {
	type args struct {
		grid     [][]int
		moveCost [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				grid:     [][]int{{5, 3}, {4, 0}, {2, 1}},
				moveCost: [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}},
			},
			want: 17,
		},
		{
			name: "Example2",
			args: args{
				grid:     [][]int{{5, 1, 2}, {4, 0, 3}},
				moveCost: [][]int{{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2}},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathCost(tt.args.grid, tt.args.moveCost); got != tt.want {
				t.Errorf("minPathCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
