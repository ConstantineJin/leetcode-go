package main

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "Example1",
			args: args{
				order: "cba",
				s:     "abcd",
			},
			wantAns: "cbad",
		},
		{
			name: "Example2",
			args: args{
				order: "bcafg",
				s:     "abcd",
			},
			wantAns: "bcad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := customSortString(tt.args.order, tt.args.s); gotAns != tt.wantAns {
				t.Errorf("customSortString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
