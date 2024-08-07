package main

import (
	"reflect"
	"testing"
)

func Test_canEat(t *testing.T) {
	type args struct {
		candiesCount []int
		queries      [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example1",
			args: args{
				candiesCount: []int{7, 4, 5, 3, 8},
				queries:      [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}},
			},
			want: []bool{true, false, true},
		},
		{
			name: "Example2",
			args: args{
				candiesCount: []int{5, 2, 6, 4, 1},
				queries:      [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}},
			},
			want: []bool{false, true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEat(tt.args.candiesCount, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}
