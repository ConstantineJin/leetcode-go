package main

import "testing"

func Test_countLargestGroup(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{n: 13},
			wantAns: 4,
		},
		{
			name:    "Example2",
			args:    args{n: 2},
			wantAns: 2,
		},
		{
			name:    "Example3",
			args:    args{n: 15},
			wantAns: 6,
		},
		{
			name:    "Example4",
			args:    args{n: 24},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countLargestGroup(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("countLargestGroup() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
