package main

import (
	"reflect"
	"testing"
)

func Test_findProductsOfElements(t *testing.T) {
	type args struct {
		queries [][]int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{queries: [][]int64{{1, 3, 7}}},
			want: []int{4},
		},
		{
			name: "Example2",
			args: args{queries: [][]int64{{2, 5, 3}, {7, 7, 4}}},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findProductsOfElements(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findProductsOfElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
