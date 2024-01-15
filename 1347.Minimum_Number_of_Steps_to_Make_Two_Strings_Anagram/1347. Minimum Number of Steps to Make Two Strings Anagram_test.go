package main

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				s: "bab",
				t: "aba",
			},
			wantAns: 1,
		},
		{
			name: "Example2",
			args: args{
				s: "leetcode",
				t: "practice",
			},
			wantAns: 5,
		},
		{
			name: "Example3",
			args: args{
				s: "anagram",
				t: "mangaar",
			},
			wantAns: 0,
		},
		{
			name: "Example4",
			args: args{
				s: "xxyyzz",
				t: "xxyyzz",
			},
			wantAns: 0,
		},
		{
			name: "Example5",
			args: args{
				s: "friend",
				t: "family",
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minSteps(tt.args.s, tt.args.t); gotAns != tt.wantAns {
				t.Errorf("minSteps() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
