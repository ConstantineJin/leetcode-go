package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{s: "abccccdd"},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{s: "a"},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{s: "aaaaaccc"},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestPalindrome(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("longestPalindrome() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
