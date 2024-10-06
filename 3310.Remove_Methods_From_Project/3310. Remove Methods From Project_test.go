package main

import (
	"reflect"
	"testing"
)

func Test_remainingMethods(t *testing.T) {
	type args struct {
		n           int
		k           int
		invocations [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				n:           4,
				k:           1,
				invocations: [][]int{{1, 2}, {0, 1}, {3, 2}},
			},
			wantAns: []int{0, 1, 2, 3},
		},
		{
			name: "Example2",
			args: args{
				n:           5,
				k:           0,
				invocations: [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}},
			},
			wantAns: []int{3, 4},
		},
		{
			name: "Example3",
			args: args{
				n:           3,
				k:           2,
				invocations: [][]int{{1, 2}, {0, 1}, {2, 0}},
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := remainingMethods(tt.args.n, tt.args.k, tt.args.invocations); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("remainingMethods() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
