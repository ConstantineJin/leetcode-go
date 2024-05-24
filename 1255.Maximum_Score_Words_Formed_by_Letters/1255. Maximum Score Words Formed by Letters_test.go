package main

import "testing"

func Test_maxScoreWords(t *testing.T) {
	type args struct {
		words   []string
		letters []byte
		score   []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				words:   []string{"dog", "cat", "dad", "good"},
				letters: []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
				score:   []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			wantAns: 23,
		},
		{
			name: "Example2",
			args: args{
				words:   []string{"xxxz", "ax", "bx", "cx"},
				letters: []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
				score:   []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10},
			},
			wantAns: 27,
		},
		{
			name: "Example3",
			args: args{
				words:   []string{"leetcode"},
				letters: []byte{'l', 'e', 't', 'c', 'o', 'd'},
				score:   []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxScoreWords(tt.args.words, tt.args.letters, tt.args.score); gotAns != tt.wantAns {
				t.Errorf("maxScoreWords() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
