package main

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			wantAns: []int{1},
		},
		{
			name: "Example2",
			args: args{
				n:     6,
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			},
			wantAns: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findMinHeightTrees() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
