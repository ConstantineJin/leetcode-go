package main

import "testing"

func Test_longestAwesome(t *testing.T) {
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
			args:    args{s: "3242415"},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{s: "12345678"},
			wantAns: 1,
		},
		{
			name:    "Example3",
			args:    args{s: "213123"},
			wantAns: 6,
		},
		{
			name:    "Example4",
			args:    args{s: "00"},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestAwesome(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("longestAwesome() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
