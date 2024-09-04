package main

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				commands:  []int{4, -1, 3},
				obstacles: nil,
			},
			wantAns: 25,
		},
		{
			name: "Example2",
			args: args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			wantAns: 65,
		},
		{
			name: "Example3",
			args: args{
				commands:  []int{6, -1, -1, 6},
				obstacles: nil,
			},
			wantAns: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := robotSim(tt.args.commands, tt.args.obstacles); gotAns != tt.wantAns {
				t.Errorf("robotSim() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
