package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: "Example2",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "Example3",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
