package main

import "testing"

func Test_maxmiumScore(t *testing.T) {
	type args struct {
		cards []int
		cnt   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				cards: []int{1, 2, 8, 9},
				cnt:   3,
			},
			wantAns: 18,
		},
		{
			name: "Example2",
			args: args{
				cards: []int{3, 3, 1},
				cnt:   1,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxmiumScore(tt.args.cards, tt.args.cnt); gotAns != tt.wantAns {
				t.Errorf("maxmiumScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
