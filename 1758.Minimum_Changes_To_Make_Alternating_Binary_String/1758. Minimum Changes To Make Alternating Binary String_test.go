package main

import "testing"

func Test_minOperations(t *testing.T) {
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
			args:    args{s: "0100"},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{s: "10"},
			wantAns: 0,
		},
		{
			name:    "Example2",
			args:    args{s: "1111"},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minOperations(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
