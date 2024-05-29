package main

import "testing"

func Test_maximumLength(t *testing.T) {
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
			args:    args{s: "aaaa"},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{s: "abcdef"},
			wantAns: -1,
		},
		{
			name:    "Example3",
			args:    args{s: "abcaba"},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumLength(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maximumLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
