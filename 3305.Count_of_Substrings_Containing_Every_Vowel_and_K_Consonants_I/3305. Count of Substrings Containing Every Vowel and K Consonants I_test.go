package main

import "testing"

func Test_countOfSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				word: "aeioqq",
				k:    1,
			},
			want: 0,
		},
		{
			name: "Example2",
			args: args{
				word: "aeiou",
				k:    0,
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				word: "ieaouqqieaouqq",
				k:    1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countOfSubstrings(tt.args.word, tt.args.k); gotAns != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
