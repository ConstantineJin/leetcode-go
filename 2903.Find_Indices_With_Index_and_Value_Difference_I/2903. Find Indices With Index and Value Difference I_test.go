package main

import (
	"reflect"
	"testing"
)

func Test_findIndices(t *testing.T) {
	type args struct {
		nums            []int
		indexDifference int
		valueDifference int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:            []int{5, 1, 4, 1},
				indexDifference: 2,
				valueDifference: 4,
			},
			want: []int{0, 3},
		},
		{
			name: "Example2",
			args: args{
				nums:            []int{2, 1},
				indexDifference: 0,
				valueDifference: 0,
			},
			want: []int{0, 0},
		},
		{
			name: "Example3",
			args: args{
				nums:            []int{1, 2, 3},
				indexDifference: 2,
				valueDifference: 4,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndices(tt.args.nums, tt.args.indexDifference, tt.args.valueDifference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
