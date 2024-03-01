package main

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{x: 4},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{x: 8},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := mySqrt(tt.args.x); gotAns != tt.wantAns {
				t.Errorf("mySqrt() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
