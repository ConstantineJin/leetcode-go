package main

import "testing"

func Test_stringHash(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "Example1",
			args: args{
				s: "abcd",
				k: 2,
			},
			wantAns: "bf",
		},
		{
			name: "Example2",
			args: args{
				s: "mxz",
				k: 3,
			},
			wantAns: "i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := stringHash(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("stringHash() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
