package main

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name:    "Example1",
			args:    args{n: 1},
			wantAns: "1",
		},
		{
			name:    "Example2",
			args:    args{n: 4},
			wantAns: "1211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countAndSay(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("countAndSay() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
