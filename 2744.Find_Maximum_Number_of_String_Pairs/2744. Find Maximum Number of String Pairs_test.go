package main

import "testing"

func Test_maximumNumberOfStringPairs(t *testing.T) {
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
			args:    args{words: []string{"cd", "ac", "dc", "ca", "zz"}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{words: []string{"ab", "ba", "cc"}},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{words: []string{"aa", "ab"}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumNumberOfStringPairs(tt.args.words); gotAns != tt.wantAns {
				t.Errorf("maximumNumberOfStringPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
