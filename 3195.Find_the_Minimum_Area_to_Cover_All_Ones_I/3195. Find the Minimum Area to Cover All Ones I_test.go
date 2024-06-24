package main

import "testing"

func Test_minimumArea(t *testing.T) {
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
			args: args{grid: [][]int{{0, 1, 0}, {1, 0, 1}}},
			want: 6,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{0, 0}, {1, 0}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumArea(tt.args.grid); got != tt.want {
				t.Errorf("minimumArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
