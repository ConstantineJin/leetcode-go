package main

import (
	"reflect"
	"testing"
)

func Test_constructGridLayout(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			},
			want: [][]int{{3, 1}, {2, 0}},
		},
		{
			name: "Example2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 3}, {2, 3}, {2, 4}},
			},
			want: [][]int{{4}, {2}, {3}, {1}, {0}},
		},
		{
			name: "Example3",
			args: args{
				n:     9,
				edges: [][]int{{0, 1}, {0, 4}, {0, 5}, {1, 7}, {2, 3}, {2, 4}, {2, 5}, {3, 6}, {4, 6}, {4, 7}, {6, 8}, {7, 8}},
			},
			want: [][]int{{8, 6, 3}, {7, 4, 2}, {1, 0, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructGridLayout(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructGridLayout() = %v, want %v", got, tt.want)
			}
		})
	}
}
