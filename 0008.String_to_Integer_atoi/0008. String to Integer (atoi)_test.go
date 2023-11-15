package main

import "testing"

func Test_myAtoi(t *testing.T) {
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
			args:    args{s: "42"},
			wantAns: 42,
		},
		{
			name:    "Example2",
			args:    args{s: "   -42"},
			wantAns: -42,
		},
		{
			name:    "Example3",
			args:    args{s: "4193 with words"},
			wantAns: 4193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := myAtoi(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("myAtoi() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
