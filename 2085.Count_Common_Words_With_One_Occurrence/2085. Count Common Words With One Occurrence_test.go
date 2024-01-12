package main

import "testing"

func Test_countWords(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				words1: []string{"leetcode", "is", "amazing", "as", "is"},
				words2: []string{"amazing", "leetcode", "is"},
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				words1: []string{"b", "bb", "bbb"},
				words2: []string{"a", "aa", "aaa"},
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				words1: []string{"a", "ab"},
				words2: []string{"a", "a", "a", "ab"},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countWords(tt.args.words1, tt.args.words2); gotAns != tt.wantAns {
				t.Errorf("countWords() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
