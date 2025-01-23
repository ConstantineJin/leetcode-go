package main

import "testing"

func Test_gridGame(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{grid: [][]int{{2, 5, 4}, {1, 5, 1}}},
			want: 4,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{3, 3, 1}, {8, 5, 2}}},
			want: 4,
		},
		{
			name: "Example3",
			args: args{grid: [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridGame(tt.args.grid); got != tt.want {
				t.Errorf("gridGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
