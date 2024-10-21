package main

import "testing"

func Test_maxUniqueSplit(t *testing.T) {
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
			args:    args{s: "ababccc"},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{s: "aba"},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{s: "aa"},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxUniqueSplit(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxUniqueSplit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
