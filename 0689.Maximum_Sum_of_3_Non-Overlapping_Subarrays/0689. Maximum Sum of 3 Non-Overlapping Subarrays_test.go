package main

import (
	"reflect"
	"testing"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 1, 2, 6, 7, 5, 1},
				k:    2,
			},
			want: []int{0, 3, 5},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 2, 1, 2, 1, 2, 1, 2, 1},
				k:    2,
			},
			want: []int{0, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumOfThreeSubarrays(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumOfThreeSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
