package main

import "testing"

func Test_maxProbability(t *testing.T) {
	type args struct {
		n                  int
		edges              [][]int
		successProbability []float64
		startNode          int
		endNode            int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{
				n:                  3,
				edges:              [][]int{{0, 1}, {1, 2}, {0, 2}},
				successProbability: []float64{0.5, 0.5, 0.2},
				startNode:          0,
				endNode:            2,
			},
			want: 0.25,
		},
		{
			name: "Example2",
			args: args{
				n:                  3,
				edges:              [][]int{{0, 1}, {1, 2}, {0, 2}},
				successProbability: []float64{0.5, 0.5, 0.3},
				startNode:          0,
				endNode:            2,
			},
			want: 0.3,
		},
		{
			name: "Example3",
			args: args{
				n:                  3,
				edges:              [][]int{{0, 1}},
				successProbability: []float64{0.5},
				startNode:          0,
				endNode:            2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProbability(tt.args.n, tt.args.edges, tt.args.successProbability, tt.args.startNode, tt.args.endNode); got != tt.want {
				t.Errorf("maxProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
