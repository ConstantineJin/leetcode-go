package main

import "testing"

func Test_maxProduct(t *testing.T) {
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
			args:    args{words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			wantAns: 16,
		},
		{
			name:    "Example2",
			args:    args{words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			wantAns: 4,
		},
		{
			name:    "Example3",
			args:    args{words: []string{"a", "aa", "aaa", "aaaa"}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxProduct(tt.args.words); gotAns != tt.wantAns {
				t.Errorf("maxProduct() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
