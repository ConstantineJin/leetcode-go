package main

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			wantAns: 6,
		},
		{
			name: "Example2",
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			wantAns: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := timeRequiredToBuy(tt.args.tickets, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("timeRequiredToBuy() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
