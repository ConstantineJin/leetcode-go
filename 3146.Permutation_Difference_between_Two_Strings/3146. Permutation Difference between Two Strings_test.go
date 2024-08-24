package main

import "testing"

func Test_findPermutationDifference(t *testing.T) {
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
				s: "abc",
				t: "bac",
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				s: "abcde",
				t: "edbac",
			},
			wantAns: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findPermutationDifference(tt.args.s, tt.args.t); gotAns != tt.wantAns {
				t.Errorf("findPermutationDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
