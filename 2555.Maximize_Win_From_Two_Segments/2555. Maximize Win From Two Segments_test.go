package main

import "testing"

func Test_maximizeWin(t *testing.T) {
	type args struct {
		prizePositions []int
		k              int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				prizePositions: []int{1, 1, 2, 2, 3, 3, 5},
				k:              2,
			},
			wantAns: 7,
		},
		{
			name: "Example2",
			args: args{
				prizePositions: []int{1, 2, 3, 4},
				k:              0,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximizeWin(tt.args.prizePositions, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maximizeWin() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
