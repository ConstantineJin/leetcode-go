package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "Example2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "Example3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "Example4",
			args: args{
				s: "abc",
				p: "a***abc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
