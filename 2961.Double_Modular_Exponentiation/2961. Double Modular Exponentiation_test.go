package main

import (
	"reflect"
	"testing"
)

func Test_getGoodIndices(t *testing.T) {
	type args struct {
		variables [][]int
		target    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				variables: [][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}},
				target:    2,
			},
			wantAns: []int{0, 2},
		},
		{
			name: "Example2",
			args: args{
				variables: [][]int{{39, 3, 1000, 1000}},
				target:    17,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getGoodIndices(tt.args.variables, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getGoodIndices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
