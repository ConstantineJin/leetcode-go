package main

import "testing"

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				numBottles:  9,
				numExchange: 3,
			},
			wantAns: 13,
		},
		{
			name: "Example2",
			args: args{
				numBottles:  15,
				numExchange: 4,
			},
			wantAns: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numWaterBottles(tt.args.numBottles, tt.args.numExchange); gotAns != tt.wantAns {
				t.Errorf("numWaterBottles() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
