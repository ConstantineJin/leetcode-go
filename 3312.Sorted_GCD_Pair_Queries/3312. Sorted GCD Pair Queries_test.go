package main

import (
	"reflect"
	"testing"
)

func Test_gcdValues(t *testing.T) {
	type args struct {
		nums    []int
		queries []int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:    []int{2, 3, 4},
				queries: []int64{0, 2, 2},
			},
			want: []int{1, 2, 2},
		},
		{
			name: "Example2",
			args: args{
				nums:    []int{4, 4, 2, 1},
				queries: []int64{5, 3, 1, 0},
			},
			want: []int{4, 2, 1, 1},
		},
		{
			name: "Example3",
			args: args{
				nums:    []int{2, 2},
				queries: []int64{0, 0},
			},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdValues(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gcdValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
