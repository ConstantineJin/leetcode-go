package main

import "testing"

func Test_numSplits(t *testing.T) {
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
			args:    args{s: "aacaba"},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{s: "abcd"},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{s: "aaaaa"},
			wantAns: 4,
		},
		{
			name:    "Example4",
			args:    args{s: "acbadbaada"},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSplits(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("numSplits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
