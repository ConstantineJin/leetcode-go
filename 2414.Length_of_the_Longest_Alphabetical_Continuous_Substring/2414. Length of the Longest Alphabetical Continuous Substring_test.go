package main

import "testing"

func Test_longestContinuousSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{s: "abacaba"},
			want: 2,
		},
		{
			name: "Example2",
			args: args{s: "abcde"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestContinuousSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestContinuousSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
