package main

import "testing"

func Test_mostBooked(t *testing.T) {
	type args struct {
		n        int
		meetings [][]int
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
				meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			},
			wantAns: 0,
		},
		{
			name: "Example2",
			args: args{
				n:        3,
				meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := mostBooked(tt.args.n, tt.args.meetings); gotAns != tt.wantAns {
				t.Errorf("mostBooked() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
