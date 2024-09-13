package main

import (
	"reflect"
	"testing"
)

func Test_getOrder(t *testing.T) {
	type args struct{ tasks [][]int }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{tasks: [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}},
			want: []int{0, 2, 3, 1},
		},
		{
			name: "Example2",
			args: args{tasks: [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}},
			want: []int{4, 3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
