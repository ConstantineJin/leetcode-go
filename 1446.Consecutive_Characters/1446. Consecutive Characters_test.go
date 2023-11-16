package main

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{s: "leetcode"},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{s: "abbcccddddeeeeedcba"},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxPower(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxPower() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
