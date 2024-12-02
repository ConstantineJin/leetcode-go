package main

import "testing"

func Test_isPrefixOfWord(t *testing.T) {
	type args struct {
		sentence   string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				sentence:   "i love eating burger",
				searchWord: "burg",
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				sentence:   "this problem is an easy problem",
				searchWord: "pro",
			},
			want: 2,
		},
		{
			name: "Example3",
			args: args{
				sentence:   "i am tired",
				searchWord: "you",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrefixOfWord(tt.args.sentence, tt.args.searchWord); got != tt.want {
				t.Errorf("isPrefixOfWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
