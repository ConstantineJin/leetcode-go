package main

import "testing"

func Test_countPaths(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example1",
			args: args{
				n:     5,
				edges: [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}},
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				n:     6,
				edges: [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {3, 6}},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPaths(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
