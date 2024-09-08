package main

import (
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			name: "Example1",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
				k:    5,
			},
			want: []*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, nil, nil},
		},
		{
			name: "Example2",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}}}}}}}}},
				k:    3,
			},
			want: []*ListNode{{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, {Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}, {Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
