package main

import (
	"reflect"
	"testing"
)

func Test_resultsArray(t *testing.T) {
	type args struct {
		queries [][]int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				queries: [][]int{{1, 2}, {3, 4}, {2, 3}, {-3, 0}},
				k:       2,
			},
			want: []int{-1, 7, 5, 3},
		},
		{
			name: "Example2",
			args: args{
				queries: [][]int{{5, 5}, {4, 4}, {3, 3}},
				k:       1,
			},
			want: []int{10, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultsArray(tt.args.queries, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
