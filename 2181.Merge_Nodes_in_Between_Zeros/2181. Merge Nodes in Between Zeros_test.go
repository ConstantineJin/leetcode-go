package main

import (
	"reflect"
	"testing"
)

func Test_mergeNodes(t *testing.T) {
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
			args: args{head: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}}},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 11}},
		},
		{
			name: "Example2",
			args: args{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
