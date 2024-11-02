package main

import "testing"

func Test_makeFancyString(t *testing.T) {
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
			args: args{s: "leeetcode"},
			want: "leetcode",
		},
		{
			name: "Example2",
			args: args{s: "aaabaaaa"},
			want: "aabaa",
		},
		{
			name: "Example3",
			args: args{s: "aab"},
			want: "aab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFancyString(tt.args.s); got != tt.want {
				t.Errorf("makeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
