package main

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns *ListNode
	}{
		{
			name: "Example1",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				k:    2,
			},
			wantAns: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			name: "Example2",
			args: args{
				head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
				k:    4,
			},
			wantAns: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("rotateRight() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
