package main

import "testing"

func Test_maximumRobots(t *testing.T) {
	type args struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				chargeTimes:  []int{3, 6, 1, 3, 4},
				runningCosts: []int{2, 1, 3, 4, 5},
				budget:       25,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				chargeTimes:  []int{11, 12, 19},
				runningCosts: []int{10, 8, 7},
				budget:       19,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumRobots(tt.args.chargeTimes, tt.args.runningCosts, tt.args.budget); gotAns != tt.wantAns {
				t.Errorf("maximumRobots() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
