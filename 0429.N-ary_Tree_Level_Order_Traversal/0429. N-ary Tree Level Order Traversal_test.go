package main

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Example1",
			args: args{
				root: &Node{
					Val:      1,
					Children: []*Node{{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4}},
				},
			},
			wantAns: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := levelOrder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("levelOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
