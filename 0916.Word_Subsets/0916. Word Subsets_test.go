package main

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Example1",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"},
			},
			wantAns: []string{"facebook", "google", "leetcode"},
		},
		{
			name: "Example2",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"l", "e"},
			},
			wantAns: []string{"apple", "google", "leetcode"},
		},
		{
			name: "Example3",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "oo"},
			},
			wantAns: []string{"facebook", "google"},
		},
		{
			name: "Example4",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"lo", "eo"},
			},
			wantAns: []string{"google", "leetcode"},
		},
		{
			name: "Example5",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"ec", "oc", "ceo"},
			},
			wantAns: []string{"facebook", "leetcode"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := wordSubsets(tt.args.words1, tt.args.words2); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("wordSubsets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
