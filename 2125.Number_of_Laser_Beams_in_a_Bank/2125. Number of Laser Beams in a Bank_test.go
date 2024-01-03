package main

import "testing"

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{bank: []string{"011001", "000000", "010100", "001000"}},
			wantAns: 8,
		},
		{
			name:    "Example2",
			args:    args{bank: []string{"000", "111", "000"}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfBeams(tt.args.bank); gotAns != tt.wantAns {
				t.Errorf("numberOfBeams() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
