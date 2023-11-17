package main

import (
	"reflect"
	"testing"
)

func Test_maximumSumQueries(t *testing.T) {
	type args struct {
		nums1   []int
		nums2   []int
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
				nums1:   []int{4, 3, 1, 2},
				nums2:   []int{2, 4, 9, 5},
				queries: [][]int{{4, 1}, {1, 3}, {2, 5}},
			},
			want: []int{6, 10, 7},
		},
		{
			name: "Example2",
			args: args{
				nums1:   []int{3, 2, 5},
				nums2:   []int{2, 3, 4},
				queries: [][]int{{4, 4}, {3, 2}, {1, 1}},
			},
			want: []int{9, 9, 9},
		},
		{
			name: "Example3",
			args: args{
				nums1:   []int{2, 1},
				nums2:   []int{2, 3},
				queries: [][]int{{3, 3}},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSumQueries(tt.args.nums1, tt.args.nums2, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumSumQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
