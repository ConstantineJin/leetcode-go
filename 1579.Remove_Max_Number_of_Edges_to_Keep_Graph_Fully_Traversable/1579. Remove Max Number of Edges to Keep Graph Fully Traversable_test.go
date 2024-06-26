package main

import "testing"

func Test_maxNumEdgesToRemove(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:     4,
				edges: [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}},
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				n:     4,
				edges: [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}},
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				n:     4,
				edges: [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumEdgesToRemove(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("maxNumEdgesToRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
