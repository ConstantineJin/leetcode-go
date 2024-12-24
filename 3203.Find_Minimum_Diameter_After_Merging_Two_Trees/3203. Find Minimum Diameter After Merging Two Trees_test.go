package main

import "testing"

func Test_minimumDiameterAfterMerge(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {0, 3}},
				edges2: [][]int{{0, 1}},
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}},
				edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDiameterAfterMerge(tt.args.edges1, tt.args.edges2); got != tt.want {
				t.Errorf("minimumDiameterAfterMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
