package main

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{n: "123"},
			want: "121",
		},
		{
			name: "Example2",
			args: args{n: "1"},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
