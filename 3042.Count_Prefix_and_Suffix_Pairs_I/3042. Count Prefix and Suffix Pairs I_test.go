package main

import "testing"

func Test_countPrefixSuffixPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{words: []string{"a", "aba", "ababa", "aa"}},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{words: []string{"pa", "papa", "ma", "mama"}},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{words: []string{"abab", "ab"}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPrefixSuffixPairs(tt.args.words); gotAns != tt.wantAns {
				t.Errorf("countPrefixSuffixPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
