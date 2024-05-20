package main

import (
	"reflect"
	"testing"
)

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example1",
			args: args{
				nums:    []int{3, 4, 1, 2, 6},
				queries: [][]int{{0, 4}},
			},
			want: []bool{false},
		},
		{
			name: "Example2",
			args: args{
				nums:    []int{4, 3, 1, 6},
				queries: [][]int{{0, 2}, {2, 3}},
			},
			want: []bool{false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
