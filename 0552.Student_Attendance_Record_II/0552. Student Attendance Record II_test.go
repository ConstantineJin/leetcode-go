package main

import "testing"

func Test_checkRecord(t *testing.T) {
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
			args:    args{n: 2},
			wantAns: 8,
		},
		{
			name:    "Example2",
			args:    args{n: 1},
			wantAns: 3,
		},
		{
			name:    "Example3",
			args:    args{n: 10101},
			wantAns: 183236316,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := checkRecord(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("checkRecord() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
