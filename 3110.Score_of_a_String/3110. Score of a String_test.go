package main

import "testing"

func Test_scoreOfString(t *testing.T) {
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
			args:    args{s: "hello"},
			wantAns: 13,
		},
		{
			name:    "Example2",
			args:    args{s: "zaz"},
			wantAns: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := scoreOfString(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("scoreOfString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
