package main

import (
	"reflect"
	"testing"
)

func Test_timeTaken(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{edges: [][]int{{0, 1}, {0, 2}}},
			want: []int{2, 4, 3},
		},
		{
			name: "Example3",
			args: args{edges: [][]int{{0, 1}}},
			want: []int{1, 2},
		},
		{
			name: "Example3",
			args: args{edges: [][]int{{2, 4}, {0, 1}, {2, 3}, {0, 2}}},
			want: []int{4, 6, 3, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeTaken(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeTaken() = %v, want %v", got, tt.want)
			}
		})
	}
}
