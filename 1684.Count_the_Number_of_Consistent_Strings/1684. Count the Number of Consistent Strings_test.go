package main

import "testing"

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				allowed: "ab",
				words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				allowed: "abc",
				words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			},
			wantAns: 7,
		},
		{
			name: "Example3",
			args: args{
				allowed: "cad",
				words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countConsistentStrings(tt.args.allowed, tt.args.words); gotAns != tt.wantAns {
				t.Errorf("countConsistentStrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
