package main

import "testing"

func Test_isAcronym(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				words: []string{"alice", "bob", "charlie"},
				s:     "abc",
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				words: []string{"an", "apple"},
				s:     "a",
			},
			want: false,
		},
		{
			name: "Example3",
			args: args{
				words: []string{"never", "gonna", "give", "up", "on", "you"},
				s:     "ngguoy",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAcronym(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("isAcronym() = %v, want %v", got, tt.want)
			}
		})
	}
}
