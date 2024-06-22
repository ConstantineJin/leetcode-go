package main

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{time: []int{30, 20, 150, 100, 40}},
			wantAns: 3,
		},
		{
			name:    "Example2",
			args:    args{time: []int{60, 60, 60}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numPairsDivisibleBy60(tt.args.time); gotAns != tt.wantAns {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
