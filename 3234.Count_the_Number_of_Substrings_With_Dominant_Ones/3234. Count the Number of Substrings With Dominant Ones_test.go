package main

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
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
			args:    args{s: "00011"},
			wantAns: 5,
		},
		{
			name:    "Example2",
			args:    args{s: "101101"},
			wantAns: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfSubstrings(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("numberOfSubstrings() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
