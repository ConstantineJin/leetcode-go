package main

import "testing"

func Test_reportSpam(t *testing.T) {
	type args struct {
		message     []string
		bannedWords []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				message:     []string{"hello", "world", "leetcode"},
				bannedWords: []string{"world", "hello"},
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				message:     []string{"hello", "programming", "fun"},
				bannedWords: []string{"world", "programming", "leetcode"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reportSpam(tt.args.message, tt.args.bannedWords); got != tt.want {
				t.Errorf("reportSpam() = %v, want %v", got, tt.want)
			}
		})
	}
}
