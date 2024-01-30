package main

import "testing"

func Test_countKeyChanges(t *testing.T) {
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
			args:    args{s: "aAbBcC"},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{s: "AaAaAaaA"},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countKeyChanges(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("countKeyChanges() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
