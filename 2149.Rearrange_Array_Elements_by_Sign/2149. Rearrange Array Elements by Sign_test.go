package main

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{nums: []int{3, 1, -2, -5, 2, -4}},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			name: "Example2",
			args: args{nums: []int{-1, 1}},
			want: []int{1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
