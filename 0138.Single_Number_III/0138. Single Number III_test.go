package main

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 2, 1, 3, 2, 5}},
			want: []int{5, 3},
		},
		{
			name: "Example 2",
			args: args{nums: []int{-1, 0}},
			want: []int{0, -1},
		},
		{
			name: "Example 3",
			args: args{nums: []int{0, 1}},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
