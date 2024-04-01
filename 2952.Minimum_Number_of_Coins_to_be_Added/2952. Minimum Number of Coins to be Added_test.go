package main

import "testing"

func Test_minimumAddedCoins(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				coins:  []int{1, 4, 10},
				target: 19,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				coins:  []int{1, 4, 10, 5, 7, 19},
				target: 19,
			},
			wantAns: 1,
		},
		{
			name: "Example3",
			args: args{
				coins:  []int{1, 1, 1},
				target: 20,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumAddedCoins(tt.args.coins, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("minimumAddedCoins() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
