package main

import (
	"reflect"
	"testing"
)

func Test_findIntersectionValues(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums1: []int{4, 3, 2, 3, 1},
				nums2: []int{2, 2, 5, 2, 3, 6},
			},
			want: []int{3, 4},
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{3, 4, 2, 3},
				nums2: []int{1, 5},
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntersectionValues(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIntersectionValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
