package main

import "testing"

func Test_maxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				difficulty: []int{2, 4, 6, 8, 10},
				profit:     []int{10, 20, 30, 40, 50},
				worker:     []int{4, 5, 6, 7},
			},
			wantAns: 100,
		},
		{
			name: "Example2",
			args: args{
				difficulty: []int{85, 47, 57},
				profit:     []int{24, 66, 99},
				worker:     []int{40, 25, 25},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxProfitAssignment(tt.args.difficulty, tt.args.profit, tt.args.worker); gotAns != tt.wantAns {
				t.Errorf("maxProfitAssignment() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
