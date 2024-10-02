package main

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
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
			args:    args{s: "())"},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{s: "((("},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minAddToMakeValid(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minAddToMakeValid() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
