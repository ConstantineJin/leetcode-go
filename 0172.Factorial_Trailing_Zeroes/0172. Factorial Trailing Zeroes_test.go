package main

import "testing"

func Test_trailingZeroes(t *testing.T) {
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
			args:    args{n: 3},
			wantAns: 0,
		},
		{
			name:    "Example2",
			args:    args{n: 5},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{n: 0},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trailingZeroes(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("trailingZeroes() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
