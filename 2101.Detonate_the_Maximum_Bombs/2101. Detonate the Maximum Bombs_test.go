package main

import "testing"

func Test_maximumDetonation(t *testing.T) {
	type args struct {
		bombs [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{bombs: [][]int{{2, 1, 3}, {6, 1, 4}}},
			wantAns: 2,
		},
		{
			name:    "Example2",
			args:    args{bombs: [][]int{{1, 1, 5}, {10, 10, 5}}},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumDetonation(tt.args.bombs); gotAns != tt.wantAns {
				t.Errorf("maximumDetonation() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
