package main

import "testing"

func Test_frogPosition(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		t      int
		target int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{
				n:      7,
				edges:  [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
				t:      2,
				target: 4,
			},
			want: float64(1) / float64(6),
		},
		{
			name: "Example2",
			args: args{
				n:      7,
				edges:  [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}},
				t:      1,
				target: 7,
			},
			want: float64(1) / float64(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frogPosition(tt.args.n, tt.args.edges, tt.args.t, tt.args.target); got != tt.want {
				t.Errorf("frogPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
