package main

import (
	"reflect"
	"testing"
)

func example1() *Node {
	node7 := Node{Val: 7, Left: nil, Right: nil, Next: nil}
	node5 := Node{Val: 5, Left: nil, Right: nil, Next: nil}
	node4 := Node{Val: 4, Left: nil, Right: nil, Next: nil}
	node3 := Node{Val: 3, Left: nil, Right: &node7, Next: nil}
	node2 := Node{Val: 2, Left: &node4, Right: &node5, Next: nil}
	node1 := Node{Val: 1, Left: &node2, Right: &node3, Next: nil}
	return &node1
}

func want1() *Node {
	node7 := Node{Val: 7, Left: nil, Right: nil, Next: nil}
	node5 := Node{Val: 5, Left: nil, Right: nil, Next: &node7}
	node4 := Node{Val: 4, Left: nil, Right: nil, Next: &node5}
	node3 := Node{Val: 3, Left: nil, Right: &node7, Next: nil}
	node2 := Node{Val: 2, Left: &node4, Right: &node5, Next: &node3}
	node1 := Node{Val: 1, Left: &node2, Right: &node3, Next: nil}
	return &node1
}

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Example1",
			args: args{example1()},
			want: want1(),
		},
		{
			name: "Example2",
			args: args{root: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
