package main

import (
	"reflect"
	"testing"
)

func Test_findColumnWidth(t *testing.T) {
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
			args: args{
				grid: [][]int{{1}, {22}, {333}},
			},
			want: []int{3},
		},
		{
			name: "Example2",
			args: args{
				grid: [][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}},
			},
			want: []int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findColumnWidth(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findColumnWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}
