package main

import "testing"

func Test_minOperationsMaxProfit(t *testing.T) {
	type args struct {
		customers    []int
		boardingCost int
		runningCost  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				customers:    []int{8, 3},
				boardingCost: 5,
				runningCost:  6,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				customers:    []int{10, 9, 6},
				boardingCost: 6,
				runningCost:  4,
			},
			wantAns: 7,
		},
		{
			name: "Example3",
			args: args{
				customers:    []int{3, 4, 0, 5, 1},
				boardingCost: 1,
				runningCost:  92,
			},
			wantAns: -1,
		},
		{
			name: "Example4",
			args: args{
				customers:    []int{0, 43, 37, 9, 23, 35, 18, 7, 45, 3, 8, 24, 1, 6, 37, 2, 38, 15, 1, 14, 39, 27, 4, 25, 27, 33, 43, 8, 44, 30, 38, 40, 20, 5, 17, 27, 43, 11, 6, 2, 30, 49, 30, 25, 32, 3, 18, 23, 45, 43, 30, 14, 41, 17, 42, 42, 44, 38, 18, 26, 32, 48, 37, 5, 37, 21, 2, 9, 48, 48, 40, 45, 25, 30, 49, 41, 4, 48, 40, 29, 23, 17, 7, 5, 44, 23, 43, 9, 35, 26, 44, 3, 26, 16, 31, 11, 9, 4, 28, 49, 43, 39, 9, 39, 37, 7, 6, 7, 16, 1, 30, 2, 4, 43, 23, 16, 39, 5, 30, 23, 39, 29, 31, 26, 35, 15, 5, 11, 45, 44, 45, 43, 4, 24, 40, 7, 36, 10, 10, 18, 6, 20, 13, 11, 20, 3, 32, 49, 34, 41, 13, 11, 3, 13, 0, 13, 44, 48, 43, 23, 12, 23, 2},
				boardingCost: 43,
				runningCost:  54,
			},
			wantAns: 993,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minOperationsMaxProfit(tt.args.customers, tt.args.boardingCost, tt.args.runningCost); gotAns != tt.wantAns {
				t.Errorf("minOperationsMaxProfit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
