package main

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		n int
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				n: 1,
				x: 2,
				y: 3,
			},
			wantAns: 6,
		},
		{
			name: "Example2",
			args: args{
				n: 5,
				x: 2,
				y: 1,
			},
			wantAns: 32,
		},
		{
			name: "Example3",
			args: args{
				n: 3,
				x: 3,
				y: 4,
			},
			wantAns: 684,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfWays(tt.args.n, tt.args.x, tt.args.y); gotAns != tt.wantAns {
				t.Errorf("numberOfWays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
