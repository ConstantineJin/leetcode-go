package main

import "testing"

func Test_countKConstraintSubstrings(t *testing.T) {
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
				s: "10101",
				k: 1,
			},
			wantAns: 12,
		},
		{
			name: "Example2",
			args: args{
				s: "1010101",
				k: 2,
			},
			wantAns: 25,
		},
		{
			name: "Example3",
			args: args{
				s: "11111",
				k: 1,
			},
			wantAns: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countKConstraintSubstrings(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countKConstraintSubstrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
