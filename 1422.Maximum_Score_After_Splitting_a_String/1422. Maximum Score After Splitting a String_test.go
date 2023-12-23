package main

import "testing"

func Test_maxScore(t *testing.T) {
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
			args:    args{s: "011101"},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{s: "00111"},
			wantAns: 5,
		},
		{
			name:    "Example3",
			args:    args{s: "1111"},
			wantAns: 3,
		},
		{
			name:    "Example4",
			args:    args{s: "00"},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxScore(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
