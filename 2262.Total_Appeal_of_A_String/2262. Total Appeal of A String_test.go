package main

import "testing"

func Test_appealSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{s: "abbca"},
			wantAns: 28,
		},
		{
			name:    "Example2",
			args:    args{s: "code"},
			wantAns: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := appealSum(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("appealSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
