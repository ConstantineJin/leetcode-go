package main

import (
	"reflect"
	"testing"
)

func Test_largestLocal(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{grid: [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}},
			want: [][]int{{9, 9}, {8, 6}},
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}},
			want: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestLocal(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
