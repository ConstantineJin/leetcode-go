package main

import "testing"

func Test_maximumPoints(t *testing.T) {
	type args struct {
		edges [][]int
		coins []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}},
				coins: []int{10, 10, 3, 3},
				k:     5,
			},
			want: 11,
		},
		{
			name: "Example2",
			args: args{
				edges: [][]int{{0, 1}, {0, 2}},
				coins: []int{8, 4, 4},
				k:     0,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPoints(tt.args.edges, tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("maximumPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
