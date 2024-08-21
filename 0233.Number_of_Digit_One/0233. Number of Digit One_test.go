package main

import "testing"

func Test_countDigitOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{n: 13},
			wantAns: 6,
		},
		{
			name:    "Example2",
			args:    args{n: 0},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countDigitOne(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("countDigitOne() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
