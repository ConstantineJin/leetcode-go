package main

import "testing"

func Test_maxMoves(t *testing.T) {
	type args struct {
		kx        int
		ky        int
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				kx:        1,
				ky:        1,
				positions: [][]int{{0, 0}},
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				kx:        0,
				ky:        2,
				positions: [][]int{{1, 1}, {2, 2}, {3, 3}},
			},
			want: 8,
		},
		{
			name: "Example3",
			args: args{
				kx:        0,
				ky:        0,
				positions: [][]int{{1, 2}, {2, 4}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMoves(tt.args.kx, tt.args.ky, tt.args.positions); got != tt.want {
				t.Errorf("maxMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
