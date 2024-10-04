package main

import "testing"

func Test_dividePlayers(t *testing.T) {
	type args struct {
		skill []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example1",
			args:    args{skill: []int{3, 2, 5, 1, 3, 4}},
			wantAns: 22,
		},
		{
			name:    "Example2",
			args:    args{skill: []int{3, 4}},
			wantAns: 12,
		},
		{
			name:    "Example3",
			args:    args{skill: []int{1, 1, 2, 3}},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := dividePlayers(tt.args.skill); gotAns != tt.wantAns {
				t.Errorf("dividePlayers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
