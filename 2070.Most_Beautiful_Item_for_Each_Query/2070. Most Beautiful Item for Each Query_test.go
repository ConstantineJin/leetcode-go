package main

import (
	"reflect"
	"testing"
)

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		items   [][]int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				items:   [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}},
				queries: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{2, 4, 5, 5, 6, 6},
		},
		{
			name: "Example2",
			args: args{
				items:   [][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}},
				queries: []int{1},
			},
			want: []int{4},
		},
		{
			name: "Example3",
			args: args{
				items:   [][]int{{10, 1000}},
				queries: []int{5},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.items, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
