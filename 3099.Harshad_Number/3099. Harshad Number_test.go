package main

import "testing"

func Test_sumOfTheDigitsOfHarshadNumber(t *testing.T) {
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
			args:    args{x: 18},
			wantAns: 9,
		},
		{
			name:    "Example2",
			args:    args{x: 23},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumOfTheDigitsOfHarshadNumber(tt.args.x); gotAns != tt.wantAns {
				t.Errorf("sumOfTheDigitsOfHarshadNumber() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
