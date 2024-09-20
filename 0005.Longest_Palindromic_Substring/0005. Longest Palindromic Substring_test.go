package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{s: "babad"},
			want: []string{"bab", "aba"},
		},
		{
			name: "Example2",
			args: args{s: "cbbd"},
			want: []string{"bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.args.s)
			var met bool
			for _, s := range tt.want {
				if s == got {
					met = true
				}
			}
			if !met {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
