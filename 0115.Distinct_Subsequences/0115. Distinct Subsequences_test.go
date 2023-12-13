package main

import "testing"

func Test_numDistinct(t *testing.T) {
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
				s: "rabbbit",
				t: "rabbit",
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				s: "babgbag",
				t: "bag",
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numDistinct(tt.args.s, tt.args.t); gotAns != tt.wantAns {
				t.Errorf("numDistinct() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
