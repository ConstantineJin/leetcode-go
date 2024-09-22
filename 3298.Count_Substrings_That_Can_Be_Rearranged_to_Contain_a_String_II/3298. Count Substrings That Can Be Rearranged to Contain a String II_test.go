package main

import "testing"

func Test_validSubstringCount(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				word1: "bcca",
				word2: "abc",
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				word1: "abcabc",
				word2: "abc",
			},
			wantAns: 10,
		},
		{
			name: "Example3",
			args: args{
				word1: "abcabc",
				word2: "aaabc",
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := validSubstringCount(tt.args.word1, tt.args.word2); gotAns != tt.wantAns {
				t.Errorf("validSubstringCount() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
