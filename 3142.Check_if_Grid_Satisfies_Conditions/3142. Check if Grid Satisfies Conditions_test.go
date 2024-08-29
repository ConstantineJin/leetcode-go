package main

import "testing"

func Test_satisfiesConditions(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{grid: [][]int{{1, 0, 2}, {1, 0, 2}}},
			want: true,
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{1, 1, 1}, {0, 0, 0}}},
			want: false,
		},
		{
			name: "Example3",
			args: args{grid: [][]int{{1}, {2}, {3}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := satisfiesConditions(tt.args.grid); got != tt.want {
				t.Errorf("satisfiesConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}
