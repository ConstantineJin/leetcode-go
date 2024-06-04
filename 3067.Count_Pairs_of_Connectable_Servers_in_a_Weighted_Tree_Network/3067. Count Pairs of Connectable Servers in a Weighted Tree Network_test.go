package main

import (
	"reflect"
	"testing"
)

func Test_countPairsOfConnectableServers(t *testing.T) {
	type args struct {
		edges       [][]int
		signalSpeed int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				edges:       [][]int{{0, 1, 1}, {1, 2, 5}, {2, 3, 13}, {3, 4, 9}, {4, 5, 2}},
				signalSpeed: 1,
			},
			want: []int{0, 4, 6, 6, 4, 0},
		},
		{
			name: "Example2",
			args: args{
				edges:       [][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}},
				signalSpeed: 3,
			},
			want: []int{2, 0, 0, 0, 0, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairsOfConnectableServers(tt.args.edges, tt.args.signalSpeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPairsOfConnectableServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
