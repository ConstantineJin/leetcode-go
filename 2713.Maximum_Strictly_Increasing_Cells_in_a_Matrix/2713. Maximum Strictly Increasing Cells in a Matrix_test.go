package main

import "testing"

func Test_maxIncreasingCells(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{mat: [][]int{{3, 1}, {3, 4}}},
			want: 2,
		},
		{
			name: "Example2",
			args: args{mat: [][]int{{1, 1}, {1, 1}}},
			want: 1,
		},
		{
			name: "Example3",
			args: args{mat: [][]int{{3, 1, 6}, {-9, 5, 7}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreasingCells(tt.args.mat); got != tt.want {
				t.Errorf("maxIncreasingCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
