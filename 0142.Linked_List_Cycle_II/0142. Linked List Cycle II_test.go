package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	node1_0, node1_1, node1_2, node1_3 := &ListNode{Val: 3}, &ListNode{Val: 2}, &ListNode{Val: 0}, &ListNode{Val: -4}
	node1_0.Next, node1_1.Next, node1_2.Next, node1_3.Next = node1_1, node1_2, node1_3, node1_1
	node2_0, node2_1 := &ListNode{Val: 1}, &ListNode{Val: 2}
	node2_0.Next, node2_1.Next = node2_1, node2_0
	node3_0 := &ListNode{Val: 1}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args{head: node1_0},
			want: node1_1,
		},
		{
			name: "Example2",
			args: args{head: node2_0},
			want: node2_0,
		},
		{
			name: "Example3",
			args: args{head: node3_0},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
