package main

import "testing"

func Test_minimumChairs(t *testing.T) {
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
			args:    args{s: "EEEEEEE"},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{s: "ELELEEL"},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{s: "ELEELEELLL"},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumChairs(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minimumChairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
