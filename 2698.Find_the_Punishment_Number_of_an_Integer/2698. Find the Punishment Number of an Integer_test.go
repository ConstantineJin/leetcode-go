package main

import "testing"

func Test_punishmentNumber(t *testing.T) {
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
			args:    args{n: 10},
			wantAns: 182,
		},
		{
			name:    "Example2",
			args:    args{n: 37},
			wantAns: 1478,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := punishmentNumber(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("punishmentNumber() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
