package main

import "testing"

func Test_minimumLength(t *testing.T) {
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
			args:    args{s: "abaacbcbb"},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{s: "aa"},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumLength(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minimumLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
