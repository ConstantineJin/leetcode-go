package main

import (
	"reflect"
	"testing"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		n         int
		edges     [][]int
		disappear []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				n:         3,
				edges:     [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}},
				disappear: []int{1, 1, 5},
			},
			want: []int{0, -1, 4},
		},
		{
			name: "Example2",
			args: args{
				n:         3,
				edges:     [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}},
				disappear: []int{1, 3, 5},
			},
			want: []int{0, 2, 3},
		},
		{
			name: "Example3",
			args: args{
				n:         2,
				edges:     [][]int{{0, 1, 1}},
				disappear: []int{1, 1},
			},
			want: []int{0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.n, tt.args.edges, tt.args.disappear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
