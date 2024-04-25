package main

import "testing"

func Test_longestIdealString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				s: "acfgbd",
				k: 2,
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				s: "abcd",
				k: 3,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestIdealString(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("longestIdealString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
