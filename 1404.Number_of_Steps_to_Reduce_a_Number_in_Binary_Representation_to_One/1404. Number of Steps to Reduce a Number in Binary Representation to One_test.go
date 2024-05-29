package main

import "testing"

func Test_numSteps(t *testing.T) {
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
			args:    args{s: "1101"},
			wantAns: 6,
		},
		{
			name:    "Example2",
			args:    args{s: "10"},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{s: "1"},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSteps(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("numSteps() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
