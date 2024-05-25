package main

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Example1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			wantAns: []string{"cats and dog", "cat sand dog"},
		},
		{
			name: "Example2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			wantAns: []string{"pineapple pen apple", "pine applepen apple", "pine apple pen apple"},
		},
		{
			name: "Example3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("wordBreak() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
