package main

import "testing"

func Test_makeSmallestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{s: "egcfe"},
			want: "efcfe",
		},
		{
			name: "Example2",
			args: args{s: "abcd"},
			want: "abba",
		},
		{
			name: "Example3",
			args: args{s: "seven"},
			want: "neven",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSmallestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("makeSmallestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
