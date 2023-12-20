package main

import (
	"reflect"
	"testing"
)

func Test_findPeakGrid(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{mat: [][]int{{1, 4}, {3, 2}}},
			want: []int{0, 1},
		},
		{
			name: "Example2",
			args: args{mat: [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakGrid(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeakGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
