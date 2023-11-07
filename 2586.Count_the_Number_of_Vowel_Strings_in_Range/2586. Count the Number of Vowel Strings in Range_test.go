package main

import "testing"

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words []string
		left  int
		right int
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name:    "Example1",
			args:    args{words: []string{"are", "amy", "u"}, left: 0, right: 2},
			wantCnt: 2,
		},
		{
			name:    "Example2",
			args:    args{words: []string{"hey", "aeo", "mu", "ooo", "artro"}, left: 1, right: 4},
			wantCnt: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := vowelStrings(tt.args.words, tt.args.left, tt.args.right); gotCnt != tt.wantCnt {
				t.Errorf("vowelStrings() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
