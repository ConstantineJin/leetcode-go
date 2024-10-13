package main

import (
	"reflect"
	"testing"
)

func Test_minBitwiseArray(t *testing.T) {
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
			args: args{nums: []int{2, 3, 5, 7}},
			want: []int{-1, 1, 4, 3},
		},
		{
			name: "Example2",
			args: args{nums: []int{11, 13, 31}},
			want: []int{9, 12, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitwiseArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minBitwiseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
