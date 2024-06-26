package main

import (
	"reflect"
	"testing"
)

func Test_goodSubsetofBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{grid: [][]int{{0, 1, 1, 0}, {0, 0, 0, 1}, {1, 1, 1, 1}}},
			want: []int{0, 1},
		},
		{
			name: "Example2",
			args: args{grid: [][]int{{0}}},
			want: []int{0},
		},
		{
			name: "Example3",
			args: args{grid: [][]int{{1, 1, 1}, {1, 1, 1}}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodSubsetofBinaryMatrix(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goodSubsetofBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
