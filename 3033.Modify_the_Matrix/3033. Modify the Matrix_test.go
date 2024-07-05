package main

import (
	"reflect"
	"testing"
)

func Test_modifiedMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{matrix: [][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}}},
			want: [][]int{{1, 2, 9}, {4, 8, 6}, {7, 8, 9}},
		},
		{
			name: "Example2",
			args: args{matrix: [][]int{{3, -1}, {5, 2}}},
			want: [][]int{{3, 2}, {5, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
