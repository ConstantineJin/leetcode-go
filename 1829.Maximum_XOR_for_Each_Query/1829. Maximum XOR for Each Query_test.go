package main

import (
	"reflect"
	"testing"
)

func Test_getMaximumXor(t *testing.T) {
	type args struct {
		nums       []int
		maximumBit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:       []int{0, 1, 1, 3},
				maximumBit: 2,
			},
			want: []int{0, 3, 2, 3},
		},
		{
			name: "Example2",
			args: args{
				nums:       []int{2, 3, 4, 7},
				maximumBit: 3,
			},
			want: []int{5, 2, 6, 5},
		},
		{
			name: "Example3",
			args: args{
				nums:       []int{0, 1, 2, 2, 5, 7},
				maximumBit: 3,
			},
			want: []int{4, 3, 6, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumXor(tt.args.nums, tt.args.maximumBit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaximumXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
