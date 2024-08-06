package main

import "testing"

func Test_minFlips(t *testing.T) {
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
			args: args{grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}},
			want: 2,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{0, 1}, {0, 1}, {0, 0}}},
			want: 1,
		},
		{
			name: "Example3",
			args: args{grid: [][]int{{1}, {0}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.grid); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
