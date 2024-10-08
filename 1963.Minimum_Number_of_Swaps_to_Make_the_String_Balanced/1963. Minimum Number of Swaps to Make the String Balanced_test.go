package main

import "testing"

func Test_minSwaps(t *testing.T) {
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
			args:    args{s: "][]["},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{s: "]]][[["},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{s: "[]"},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minSwaps(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minSwaps() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
