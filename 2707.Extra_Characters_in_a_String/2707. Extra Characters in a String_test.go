package main

import "testing"

func Test_minExtraChar(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				s:          "leetscode",
				dictionary: []string{"leet", "code", "leetcode"},
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				s:          "sayhelloworld",
				dictionary: []string{"hello", "world"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minExtraChar(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("minExtraChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
