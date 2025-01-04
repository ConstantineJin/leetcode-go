package main

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
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
			args:    args{s: "aabca"},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{s: "adc"},
			wantAns: 0,
		},
		{
			name:    "Example3",
			args:    args{s: "bbcbaba"},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPalindromicSubsequence(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
