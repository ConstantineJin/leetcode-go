package main

import "testing"

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				routes: [][]int{{1, 2, 7}, {3, 6, 7}},
				source: 1,
				target: 6,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				routes: [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}},
				source: 15,
				target: 12,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
