package main

import (
	"reflect"
	"testing"
)

func Test_platesBetweenCandles(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				s:       "**|**|***|",
				queries: [][]int{{2, 5}, {5, 9}},
			},
			want: []int{2, 3},
		},
		{
			name: "Example2",
			args: args{
				s:       "***|**|*****|**||**|*",
				queries: [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
			},
			want: []int{9, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := platesBetweenCandles(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("platesBetweenCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
