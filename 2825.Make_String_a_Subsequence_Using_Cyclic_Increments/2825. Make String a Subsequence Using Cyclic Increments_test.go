package main

import "testing"

func Test_canMakeSubsequence(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				str1: "abc",
				str2: "ad",
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				str1: "zc",
				str2: "ad",
			},
			want: true,
		},
		{
			name: "Example3",
			args: args{
				str1: "ab",
				str2: "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeSubsequence(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("canMakeSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
