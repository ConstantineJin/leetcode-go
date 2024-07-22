package main

import "testing"

func Test_maxOperations(t *testing.T) {
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
			args:    args{s: "1001101"},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{s: "00111"},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxOperations(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
