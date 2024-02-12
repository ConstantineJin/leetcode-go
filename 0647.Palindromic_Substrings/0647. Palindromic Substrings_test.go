package main

import "testing"

func Test_countSubstrings(t *testing.T) {
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
			args:    args{s: "abc"},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{s: "aaa"},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSubstrings(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("countSubstrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
