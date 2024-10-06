package main

import "testing"

func Test_areSentencesSimilar(t *testing.T) {
	type args struct {
		sentence1 string
		sentence2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				sentence1: "My name is Haley",
				sentence2: "My Haley",
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				sentence1: "of",
				sentence2: "A lot of words",
			},
			want: false,
		},
		{
			name: "Example3",
			args: args{
				sentence1: "Eating right now",
				sentence2: "Eating",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.args.sentence1, tt.args.sentence2); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
