package main

import "testing"

func Test_getMaxRepetitions(t *testing.T) {
	type args struct {
		s1 string
		n1 int
		s2 string
		n2 int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				s1: "acb",
				n1: 4,
				s2: "ab",
				n2: 2,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				s1: "acb",
				n1: 1,
				s2: "acb",
				n2: 1,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getMaxRepetitions(tt.args.s1, tt.args.n1, tt.args.s2, tt.args.n2); gotAns != tt.wantAns {
				t.Errorf("getMaxRepetitions() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
