package main

import "testing"

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:           6,
				connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				n:           5,
				connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				n:           3,
				connections: [][]int{{1, 0}, {2, 0}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minReorder(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minReorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
