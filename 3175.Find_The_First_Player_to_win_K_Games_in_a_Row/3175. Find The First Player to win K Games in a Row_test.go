package main

import "testing"

func Test_findWinningPlayer(t *testing.T) {
	type args struct {
		skills []int
		k      int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				skills: []int{4, 2, 6, 3, 9},
				k:      2,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				skills: []int{2, 5, 4},
				k:      3,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findWinningPlayer(tt.args.skills, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("findWinningPlayer() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
