package main

import "testing"

func Test_hasCycle(t *testing.T) {
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
		want bool
	}{
		{
			name: "Example1",
			args: args{head: node1_0},
			want: true,
		},
		{
			name: "Example2",
			args: args{head: node2_0},
			want: true,
		},
		{
			name: "Example3",
			args: args{head: node3_0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
