package main

import (
	"reflect"
	"testing"
)

func Test_doubleIt(t *testing.T) {
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
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}}},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}},
		},
		{
			name: "Example2",
			args: args{head: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleIt(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doubleIt() = %v, want %v", got, tt.want)
			}
		})
	}
}
