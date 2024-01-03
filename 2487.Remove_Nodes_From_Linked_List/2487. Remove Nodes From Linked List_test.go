package main

import (
	"reflect"
	"testing"
)

func Test_removeNodes(t *testing.T) {
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
			args: args{head: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 13, Next: &ListNode{Val: 3, Next: &ListNode{Val: 8}}}}}},
			want: &ListNode{Val: 13, Next: &ListNode{Val: 8}},
		},
		{
			name: "Example2",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
