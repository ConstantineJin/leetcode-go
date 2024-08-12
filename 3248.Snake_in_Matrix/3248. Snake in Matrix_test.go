package main

import "testing"

func Test_finalPositionOfSnake(t *testing.T) {
	type args struct {
		n        int
		commands []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				n:        2,
				commands: []string{"RIGHT", "DOWN"},
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				n:        3,
				commands: []string{"DOWN", "RIGHT", "UP"},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := finalPositionOfSnake(tt.args.n, tt.args.commands); gotAns != tt.wantAns {
				t.Errorf("finalPositionOfSnake() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
