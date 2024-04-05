package main

import "testing"

func Test_maxDepth(t *testing.T) {
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
			args:    args{s: "(1+(2*3)+((8)/4))+1"},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{s: "(1)+((2))+(((3)))"},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxDepth(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxDepth() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
