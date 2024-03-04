package main

import "testing"

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		power  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				tokens: []int{100},
				power:  50,
			},
			wantAns: 0,
		},
		{
			name: "Example2",
			args: args{
				tokens: []int{200, 100},
				power:  150,
			},
			wantAns: 1,
		},
		{
			name: "Example3",
			args: args{
				tokens: []int{100, 200, 300, 400},
				power:  200,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := bagOfTokensScore(tt.args.tokens, tt.args.power); gotAns != tt.wantAns {
				t.Errorf("bagOfTokensScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
