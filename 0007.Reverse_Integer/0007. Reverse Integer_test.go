package main

import "testing"

func Test_reverse(t *testing.T) {
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
			args:    args{x: 123},
			wantAns: 321,
		},
		{
			name:    "Example2",
			args:    args{x: -123},
			wantAns: -321,
		},
		{
			name:    "Example3",
			args:    args{x: 120},
			wantAns: 21,
		},
		{
			name:    "Example4",
			args:    args{x: 0},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := reverse(tt.args.x); gotAns != tt.wantAns {
				t.Errorf("reverse() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
