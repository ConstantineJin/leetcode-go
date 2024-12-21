package main

import "testing"

func Test_maxKDivisibleComponents(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		values []int
		k      int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				n:      5,
				edges:  [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
				values: []int{1, 8, 1, 4, 4},
				k:      6,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				n:      7,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
				values: []int{3, 0, 6, 1, 5, 2, 1},
				k:      3,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxKDivisibleComponents(tt.args.n, tt.args.edges, tt.args.values, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxKDivisibleComponents() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
