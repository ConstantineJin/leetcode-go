package main

import "testing"

func Test_isSubstringPresent(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{s: "leetcode"},
			want: true,
		},
		{
			name: "Example2",
			args: args{s: "abcba"},
			want: true,
		},
		{
			name: "Example3",
			args: args{s: "abcd"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubstringPresent(tt.args.s); got != tt.want {
				t.Errorf("isSubstringPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}
