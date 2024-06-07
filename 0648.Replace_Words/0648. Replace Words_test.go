package main

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				dictionary: []string{"cat", "bat", "rat"},
				sentence:   "the cattle was rattled by the battery",
			},
			want: "the cat was rat by the bat",
		},
		{
			name: "Example2",
			args: args{
				dictionary: []string{"a", "b", "c"},
				sentence:   "aadsfasf absbs bbab cadsfafs",
			},
			want: "a a b c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
