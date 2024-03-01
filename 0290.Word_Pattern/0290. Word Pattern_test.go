package main

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{pattern: "abba", s: "dog cat cat dog"},
			want: true,
		},
		{
			name: "Example2",
			args: args{pattern: "abba", s: "dog cat cat fish"},
			want: false,
		},
		{
			name: "Example3",
			args: args{pattern: "aaaa", s: "dog cat cat dog"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
