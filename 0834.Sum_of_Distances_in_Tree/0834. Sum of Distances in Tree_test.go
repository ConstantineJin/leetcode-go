package main

import (
	"reflect"
	"testing"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			name: "Example2",
			args: args{
				n:     1,
				edges: nil,
			},
			want: []int{0},
		},
		{
			name: "Example3",
			args: args{
				n:     2,
				edges: [][]int{{1, 0}},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDistancesInTree(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
