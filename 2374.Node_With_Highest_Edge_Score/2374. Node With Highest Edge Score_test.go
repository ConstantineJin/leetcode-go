package main

import "testing"

func Test_edgeScore(t *testing.T) {
	type args struct {
		edges []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{edges: []int{1, 0, 0, 0, 0, 7, 7, 5}},
			wantAns: 7,
		},
		{
			name:    "Example2",
			args:    args{edges: []int{2, 0, 0, 2}},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := edgeScore(tt.args.edges); gotAns != tt.wantAns {
				t.Errorf("edgeScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
