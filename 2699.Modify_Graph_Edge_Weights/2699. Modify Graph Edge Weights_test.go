package main

import (
	"reflect"
	"testing"
)

func Test_modifiedGraphEdges(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
		target      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				n:           5,
				edges:       [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}},
				source:      0,
				destination: 1,
				target:      5,
			},
			want: [][]int{{4, 1, 1}, {2, 0, 3}, {0, 3, 3}, {4, 3, 1}},
		},
		{
			name: "Example2",
			args: args{
				n:           3,
				edges:       [][]int{{0, 1, -1}, {0, 2, 5}},
				source:      0,
				destination: 2,
				target:      6,
			},
			want: nil,
		},
		{
			name: "Example3",
			args: args{
				n:           4,
				edges:       [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}},
				source:      0,
				destination: 2,
				target:      6,
			},
			want: [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedGraphEdges(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedGraphEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
