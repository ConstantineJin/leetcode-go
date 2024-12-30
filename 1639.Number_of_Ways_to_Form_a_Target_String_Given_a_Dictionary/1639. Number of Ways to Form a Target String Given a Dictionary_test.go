package main

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		words  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				words:  []string{"acca", "bbbb", "caca"},
				target: "aba",
			},
			want: 6,
		},
		{
			name: "Example2",
			args: args{
				words:  []string{"abba", "baab"},
				target: "bab",
			},
			want: 4,
		},
		{
			name: "Example3",
			args: args{
				words:  []string{"abcd"},
				target: "abcd",
			},
			want: 1,
		},
		{
			name: "Example4",
			args: args{
				words:  []string{"abab", "baba", "abba", "baab"},
				target: "abba",
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.words, tt.args.target); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
