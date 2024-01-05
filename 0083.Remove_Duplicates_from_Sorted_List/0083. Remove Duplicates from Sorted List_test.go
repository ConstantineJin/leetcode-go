package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "Example2",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "Example3",
			args: args{head: &ListNode{Val: 1}},
			want: &ListNode{Val: 1},
		},
		{
			name: "Example4",
			args: args{nil},
			want: nil,
		},
		{
			name: "Example5",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}},
			want: &ListNode{Val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
