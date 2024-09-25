package main

import "testing"

func Test_distinctNames(t *testing.T) {
	type args struct {
		ideas []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{ideas: []string{"coffee", "donuts", "time", "toffee"}},
			wantAns: 6,
		},
		{
			name:    "Example2",
			args:    args{ideas: []string{"lack", "back"}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := distinctNames(tt.args.ideas); gotAns != tt.wantAns {
				t.Errorf("distinctNames() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
